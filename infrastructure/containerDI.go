package infrastructure

import (
	myLib_elastic "github.com/guilherme-souza-lima/solar-system/elastic"
	"go.mongodb.org/mongo-driver/mongo"
	"ssMercurio/infrastructure/crypto"
	"ssMercurio/infrastructure/database"
	"ssMercurio/infrastructure/database/dbmongo"
	"ssMercurio/infrastructure/jwt"
	"ssMercurio/user_case/handler"
	"ssMercurio/user_case/repository"
	"ssMercurio/user_case/service"
)

type ContainerDI struct {
	Config         Config
	MongoDB        *mongo.Client
	UserRepository repository.UserRepository
	PointsService  service.PointsService
	UserService    service.UserService
	PointsHandler  handler.PointsHandler
	UserHandler    handler.UserHandler
	JwtToken       jwt.TokenJwt
	CryptoPassword crypto.CryptoPassword
	Logger         myLib_elastic.LoggerElastic
}

func NewContainerDI(config Config) *ContainerDI {
	container := &ContainerDI{
		Config: config,
	}

	mongoConfig := database.Config{
		Hostname: container.Config.DbHostMongo,
		Port:     container.Config.DbPortMongo,
		UserName: container.Config.DbUserMongo,
		Password: container.Config.DbPasswordMongo,
		Database: "",
	}
	container.MongoDB = dbmongo.InitMongo(&mongoConfig)

	brokers := []string{container.Config.BrokerElastic1, container.Config.BrokerElastic2}
	container.Logger = myLib_elastic.NewLoggerElastic(container.Config.BrokerElasticDB, brokers)

	container.buildValidation()
	container.build()
	return container
}

func (c *ContainerDI) buildValidation() {
	c.JwtToken = jwt.NewTokenJwt(c.Config.AccessSecret)
	c.CryptoPassword = crypto.NewCryptoPassword(c.Config.AccessSecret)
}
func (c *ContainerDI) build() {
	c.UserRepository = repository.NewUserRepository(c.MongoDB)
	c.UserService = service.NewUserService(c.UserRepository, c.JwtToken, c.CryptoPassword)
	c.PointsService = service.NewPointsService(c.JwtToken, c.UserRepository)
	c.UserHandler = handler.NewUserHandler(c.UserService, c.Logger)
	c.PointsHandler = handler.NewPointsHandler(c.PointsService)
}
func (c *ContainerDI) ShutDown() {}
