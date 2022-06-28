package handler

import (
	"github.com/gofiber/fiber/v2"
	"ssMercurio/user_case/request"
	"ssMercurio/user_case/response"
)

type UserService interface {
	Create(data request.User) error
	Login(data request.Login) (response.UserLogin, error)
	Verify(data request.Verify) (bool, error)
}

type UserHandler struct {
	UserService UserService
}

func NewUserHandler(UserService UserService) UserHandler {
	return UserHandler{UserService}
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
		return c.Status(fiber.StatusBadRequest).JSON("Error body parser request. Error: " + err.Error())
	}
	if user.Login == "" || user.Password == "" {
		return c.Status(fiber.StatusNotFound).JSON("Error field not filled. Error: username or password")
	}

	result, err := u.UserService.Login(user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("Error: " + err.Error())
	}

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
