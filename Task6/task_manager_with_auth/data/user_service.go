package data

import (
	"context"
	"errors"
	"os"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	cur, err := userCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var user models.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())
	return users, nil
}

func CreateUser(user models.User) error {
	if len(user.Password) < 4 {
		return errors.New("password length must be greater than 4")
	}

	users, err := GetUsers()
	if err != nil {
		return err
	}
	if len(users) == 0 {
		user.Role = "Admin"
	} else {
		user.Role = "User"
	}

	for _, existingUser := range users {
		if user.Username == existingUser.Username {
			return errors.New("username already exists")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func Login(user models.User) (string, error) {
	if len(user.Password) < 4 {
		return "", errors.New("password length must be greater than 4")
	}

	users, err := GetUsers()
	if err != nil {
		return "", err
	}

	var prevUser models.User
	for _, existingUser := range users {
		if existingUser.Username == user.Username {
			prevUser = existingUser
			break
		}
	}

	if prevUser == (models.User{}) || bcrypt.CompareHashAndPassword([]byte(prevUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": prevUser.Username,
		"role":     prevUser.Role,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")

	jwtToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func PromoteUser(username string) (*models.User, error) {
	filter := bson.D{{Key:"username", Value: username}}
	var user models.User

	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Role == "Admin" {
		return nil, errors.New("user is already an admin")
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "role", Value: "Admin"},
		}}}

	result := userCollection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updatedUser models.User
	err = result.Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
