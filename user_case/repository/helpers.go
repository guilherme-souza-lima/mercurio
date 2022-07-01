package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"ssMercurio/entities"
)

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

// getUserByCell empty result TRUE, cellphone already registered
func (m UserRepository) getUserByCell(cellphone string) bool {
	var user entities.User
	collection := m.db.Database(entities.DATABASE).Collection(entities.COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"cellphone": cellphone}).Decode(&user)
	if result != nil {
		return false
	}
	return true
}
