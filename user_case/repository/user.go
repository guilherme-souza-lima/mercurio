package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ssMercurio/entities"
	"ssMercurio/utils"
)

type RepositoryUser interface {
	Create(data entities.User) error
	Login(login entities.Login) (user entities.User, err error)
	GetUserById(ID string) (*entities.User, error)
	UpdatePointsUser(id string, point int) error
}

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
	if m.getUserByCell(user.Cellphone) {
		return errors.New("cellphone already registered")
	}
	coll := m.db.Database(entities.DATABASE).Collection(entities.COLLECTION)
	_, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		utils.LoggerWriting("Warn", "ERROR REPOSITORY CREATE: "+err.Error())
		return err
	}
	return nil
}

func (m UserRepository) Login(login entities.Login) (user entities.User, err error) {
	collection := m.db.Database(entities.DATABASE).Collection(entities.COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"email": login.Login, "password": login.Password}).Decode(&user)
	if result != nil {
		return user, result
	}
	return user, err
}

func (m UserRepository) GetUserById(ID string) (*entities.User, error) {
	var user entities.User
	collection := m.db.Database(entities.DATABASE).Collection(entities.COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"id": ID}).Decode(&user)
	if result != nil {
		return &user, result
	}
	return &user, nil
}

func (m UserRepository) UpdatePointsUser(id string, point int) error {
	collection := m.db.Database(entities.DATABASE).Collection(entities.COLLECTION)
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": id},
		bson.M{"$set": bson.M{"points.ggpoints": point}})
	if err != nil {
		utils.LoggerWriting("Warn", "ERROR REPOSITORY UPDATE POINTS: "+err.Error())
		return err
	}
	return nil
}
