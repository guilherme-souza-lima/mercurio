package infrastructure

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Environment     string
	App             string
	DbHostMongo     string
	DbPortMongo     string
	DbUserMongo     string
	DbPasswordMongo string
	AccessSecret    string
	BrokerElastic1  string
	BrokerElastic2  string
	BrokerElasticDB string
}

func NewConfig() Config {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("Error loading env file")
		}
	}

	return Config{
		Environment:     os.Getenv("ENVIRONMENT"),
		App:             os.Getenv("APP"),
		DbHostMongo:     os.Getenv("HOST_MONGO"),
		DbPortMongo:     os.Getenv("PORT_MONGO"),
		DbUserMongo:     os.Getenv("USER_MONGO"),
		DbPasswordMongo: os.Getenv("PASSWORD_MONGO"),
		AccessSecret:    os.Getenv("ACCESS_SECRET"),
		BrokerElastic1:  os.Getenv("BROKER_ELASTIC"),
		BrokerElastic2:  os.Getenv("BROKER_ELASTIC_2"),
		BrokerElasticDB: os.Getenv("BROKER_ELASTIC_DB"),
	}
}
