package dbmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ssMercurio/infrastructure/database"
)

func InitMongo(config *database.Config) *mongo.Client {
	uri := "mongodb://" + config.UserName + ":" + config.Password + "@" +
		config.Hostname + ":" + config.Port + "/?maxPoolSize=20&w=majority"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("failed to connect mongodb database")
	}
	return client
}
