package usecases

import (
	"context"
	"errors"
	domain "task_manager/Domain"
	infrastructure "task_manager/Infrastructure"
	"time"
)

type userUseCase struct {
	userRepository   domain.UserRepository
	passwordService  infrastructure.PasswordService
	jwtService       infrastructure.JWTService
	contextTimeout    time.Duration
}

func NewUserUseCase(userRepo domain.UserRepository, passwordService infrastructure.PasswordService, jwtService infrastructure.JWTService, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository:  userRepo,
		passwordService: passwordService,
		jwtService:      jwtService,
		contextTimeout:  timeout,
	}
}

func (u *userUseCase) GetUsers(ctx context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetUsers(ctx)
}

func (u *userUseCase) CreateUser(ctx context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if len(user.Password) < 4 {
		return errors.New("password length must be greater than 4")
	}

	users, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		user.Role = "Admin"
	} else {
		user.Role = "User"
	}

	existingUser, err := u.userRepository.FindByUsername(ctx, user.Username)
	if err == nil && existingUser.Username != "" {
		return errors.New("username already exists")
	}

	hashedPassword, err := u.passwordService.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return u.userRepository.CreateUser(ctx, user)
}

func (u *userUseCase) Login(ctx context.Context, user domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if len(user.Password) < 4 {
		return "", errors.New("password length must be greater than 4")
	}

	existingUser, err := u.userRepository.FindByUsername(ctx, user.Username)
	if err != nil || existingUser.Username == "" || u.passwordService.CompareHashAndPassword(existingUser.Password, user.Password) != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := u.jwtService.GenerateToken(existingUser.Username, existingUser.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUseCase) PromoteUser(ctx context.Context, username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil || user.Username == "" {
		return nil, errors.New("user not found")
	}

	if user.Role == "Admin" {
		return nil, errors.New("user is already an admin")
	}

	return u.userRepository.PromoteUser(ctx, username)
}
