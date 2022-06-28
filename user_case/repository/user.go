package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ssMercurio/entities"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) UserRepository {
	return UserRepository{db}
}

func (m UserRepository) Create(user entities.User) error {
	if m.getUserByEmail(user.Email) {
		return errors.New("e-mail already registered")
	}
	coll := m.db.Database(entities.DATABASE).Collection(entities.COLLECTION)
	_, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

// getUserByEmail empty result TRUE, email already registered
func (m UserRepository) getUserByEmail(email string) bool {
	var user entities.User
	collection := m.db.Database(entities.DATABASE).Collection(entities.COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if result != nil {
		return false
	}
	return true
}

func (m UserRepository) Login(login entities.Login) (user entities.User, err error) {
	collection := m.db.Database(entities.DATABASE).Collection(entities.COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"email": login.Login, "password": login.Password}).Decode(&user)
	if result != nil {
		return user, result
	}
	return user, err
}
