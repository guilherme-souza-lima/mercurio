package handler

import (
	"github.com/gofiber/fiber/v2"
	"ssMercurio/user_case/request"
)

type PointsService interface {
	VerifyPoints(data request.PointsRequest) (bool, error)
}

type PointsHandler struct {
	PointsService PointsService
}

func NewPointsHandler(PointsService PointsService) PointsHandler {
	return PointsHandler{PointsService}
}

func (p PointsHandler) LoginGame(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	pointsID := c.Params("points_id")
	var verify request.PointsRequest
	if err := c.BodyParser(&verify); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error body parser request. Error: " + err.Error())
	}
	verify.Verify.ID = userID
	verify.Verify.IDPoints = pointsID
	result, err := p.PointsService.VerifyPoints(verify)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error validation. Error: " + err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
