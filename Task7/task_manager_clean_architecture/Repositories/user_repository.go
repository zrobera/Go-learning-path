package repositories

import (
	"context"
	domain "task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(database mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   database,
		collection: collection,
	}
}

func (u *userRepository) CreateUser(c context.Context, user domain.User) error {
	collection := u.database.Collection(u.collection)

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindByUsername(c context.Context, username string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	filter := bson.D{{Key: "username", Value: username}}

	var user domain.User
	err := collection.FindOne(c, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetUsers(c context.Context) ([]domain.User, error) {
	collection := u.database.Collection(u.collection)

	var users []domain.User

	cur, err := collection.Find(c, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(c) {
		var user domain.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(c)
	return users, nil
}

func (u *userRepository) PromoteUser(c context.Context, username string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)

	filter := bson.D{{Key: "username", Value: username}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "role", Value: "Admin"},
		}}}

	result := collection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updatedUser domain.User
	err := result.Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
