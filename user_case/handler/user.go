package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	myLib_elastic "github.com/guilherme-souza-lima/solar-system/elastic"
	"ssMercurio/user_case/request"
	"ssMercurio/user_case/response"
	"ssMercurio/utils"
	"strconv"
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
		strError := "Error: " + err.Error()
		u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusBadRequest,
			"Handler/user CreateUser()", strError, "not logged in"))
		return c.Status(fiber.StatusBadRequest).JSON(strError)
	}
	if err := u.UserService.Create(user); err != nil {
		strError := "Error: " + err.Error()
		u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusBadRequest,
			"Handler/user u.UserService.Create(user)", strError, "not logged in"))
		return c.Status(fiber.StatusBadRequest).JSON(strError)
	}
	str := fmt.Sprintf("%#v", user)
	u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusOK, "Success", str, "Create User Success"))
	return c.Status(fiber.StatusOK).JSON("success")
}

func (u UserHandler) LoginUser(c *fiber.Ctx) error {
	var user request.Login
	if err := c.BodyParser(&user); err != nil {
		strError := "Error body parser request. Error: " + err.Error()
		u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusBadRequest,
			"Handler/user LoginUser()", strError, "not logged in"))
		return c.Status(fiber.StatusBadRequest).JSON(strError)
	}
	if user.Login == "" || user.Password == "" {
		strError := "Error field not filled. Error: username or password"
		u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusNotFound,
			"Handler/user LoginUser()", strError, "not logged in"))
		return c.Status(fiber.StatusNotFound).JSON(strError)
	}

	result, err := u.UserService.Login(user)
	if err != nil {
		strError := "Error: " + err.Error()
		u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusNotFound,
			"Handler/user u.UserService.Login(user)", strError, "not logged in"))
		return c.Status(fiber.StatusNotFound).JSON(strError)
	}

	str := fmt.Sprintf("%#v", result)
	u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusOK, "Success", str, result.ID))

	return c.Status(fiber.StatusOK).JSON(result)
}

func (u UserHandler) VerifyUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	pointsID := c.Params("points_id")
	var verify request.Verify
	if err := c.BodyParser(&verify); err != nil {
		strError := "Error body parser request. Error: " + err.Error()
		u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusBadRequest,
			"Handler/user VerifyUser()", strError, "not verify user"))
		return c.Status(fiber.StatusBadRequest).JSON(strError)
	}

	verify.ID = userID
	verify.IDPoints = pointsID
	result, err := u.UserService.Verify(verify)

	if err != nil {
		strError := err.Error()
		u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusUnauthorized,
			"Handler/user VerifyUser()", strError, "Unauthorized"))
		return c.Status(fiber.StatusUnauthorized).JSON(strError)
	}

	str := fmt.Sprintf("%#v", verify)
	u.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusOK, strconv.FormatBool(result), str, userID))
	return c.Status(fiber.StatusOK).JSON(result)
}
