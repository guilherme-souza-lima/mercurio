package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	myLib_elastic "github.com/guilherme-souza-lima/solar-system/elastic"
	"ssMercurio/user_case/request"
	"ssMercurio/utils"
	"strconv"
)

type PointsService interface {
	VerifyPoints(data request.PointsRequest) (bool, error)
}

type PointsHandler struct {
	PointsService PointsService
	Logger        myLib_elastic.LoggerElastic
}

func NewPointsHandler(PointsService PointsService, Logger myLib_elastic.LoggerElastic) PointsHandler {
	return PointsHandler{PointsService, Logger}
}

func (p PointsHandler) LoginGame(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	pointsID := c.Params("points_id")
	var verify request.PointsRequest
	if err := c.BodyParser(&verify); err != nil {
		strError := "Error body parser request. Error: " + err.Error()
		utils.LoggerWriting("Error", strError)
		p.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusBadRequest,
			"Handler/points_gg LoginGame()", strError, userID))
		return c.Status(fiber.StatusBadRequest).JSON(strError)
	}
	verify.Verify.ID = userID
	verify.Verify.IDPoints = pointsID
	result, err := p.PointsService.VerifyPoints(verify)
	if err != nil {
		strError := "Error validation. Error: " + err.Error()
		utils.LoggerWriting("Error", strError)
		p.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusBadRequest,
			"Handler/points_gg p.PointsService.VerifyPoints(verify)", strError, userID))
		return c.Status(fiber.StatusBadRequest).JSON(strError)
	}

	utils.LoggerWriting("Info", "success")
	str := fmt.Sprintf("%#v", verify)
	p.Logger.LoggerElasticsearch(utils.MappingLoggerElastic(fiber.StatusOK, strconv.FormatBool(result), str, userID))

	return c.Status(fiber.StatusOK).JSON(result)
}
