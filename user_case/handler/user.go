package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	myLib_elastic "github.com/guilherme-souza-lima/solar-system/elastic"
	myLib_entity "github.com/guilherme-souza-lima/solar-system/entity"
	"ssMercurio/entities"
	"ssMercurio/user_case/request"
	"ssMercurio/user_case/response"
	"ssMercurio/utils"
	"time"
)

type UserService interface {
	Create(data request.User) error
	Login(data request.Login) (response.UserLogin, error)
	Verify(data request.Verify) (bool, error)
}

type UserHandler struct {
	UserService UserService
	Logger      myLib_elastic.LoggerElastic
}

func NewUserHandler(UserService UserService, Logger myLib_elastic.LoggerElastic) UserHandler {
	return UserHandler{UserService, Logger}
}

func (u UserHandler) CreateUser(c *fiber.Ctx) error {
	var user request.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error: " + err.Error())
	}
	if err := u.UserService.Create(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error: " + err.Error())
	}
	return c.Status(fiber.StatusOK).JSON("success")
}

func (u UserHandler) LoginUser(c *fiber.Ctx) error {
	var user request.Login
	if err := c.BodyParser(&user); err != nil {
		logger := myLib_entity.MappingElastic{
			StatusCode: fiber.StatusBadRequest,
			Level:      "Error body parser request",
			Message: myLib_entity.MessageElastic{
				Message: err.Error(),
				Local:   entities.NAME_SYSTEM,
			},
			Date: time.Now(),
			User: myLib_entity.UserElastic{
				ID: utils.GeneratorUUid(),
			},
		}
		u.Logger.LoggerElasticsearch(logger)
		return c.Status(fiber.StatusBadRequest).JSON("Error body parser request. Error: " + err.Error())
	}
	if user.Login == "" || user.Password == "" {
		logger := myLib_entity.MappingElastic{
			StatusCode: fiber.StatusNotFound,
			Level:      "Error field not filled",
			Message: myLib_entity.MessageElastic{
				Message: "username or password empty",
				Local:   entities.NAME_SYSTEM,
			},
			Date: time.Now(),
			User: myLib_entity.UserElastic{
				ID: utils.GeneratorUUid(),
			},
		}
		u.Logger.LoggerElasticsearch(logger)
		return c.Status(fiber.StatusNotFound).JSON("Error field not filled. Error: username or password")
	}

	result, err := u.UserService.Login(user)
	if err != nil {
		logger := myLib_entity.MappingElastic{
			StatusCode: fiber.StatusNotFound,
			Level:      "Error Login",
			Message: myLib_entity.MessageElastic{
				Message: err.Error(),
				Local:   entities.NAME_SYSTEM,
			},
			Date: time.Now(),
			User: myLib_entity.UserElastic{
				ID: utils.GeneratorUUid(),
			},
		}
		u.Logger.LoggerElasticsearch(logger)
		return c.Status(fiber.StatusNotFound).JSON("Error: " + err.Error())
	}

	str := fmt.Sprintf("%#v", result)
	logger := myLib_entity.MappingElastic{
		StatusCode: fiber.StatusOK,
		Level:      "Success",
		Message: myLib_entity.MessageElastic{
			Message: str,
			Local:   entities.NAME_SYSTEM,
		},
		Date: time.Now(),
		User: myLib_entity.UserElastic{
			ID: utils.GeneratorUUid(),
		},
	}
	u.Logger.LoggerElasticsearch(logger)

	return c.Status(fiber.StatusOK).JSON(result)
}

func (u UserHandler) VerifyUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	pointsID := c.Params("points_id")
	var verify request.Verify
	if err := c.BodyParser(&verify); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error body parser request. Error: " + err.Error())
	}

	verify.ID = userID
	verify.IDPoints = pointsID
	result, err := u.UserService.Verify(verify)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
