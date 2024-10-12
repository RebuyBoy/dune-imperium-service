package handlers

import (
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PlayerHandler interface {
	Create(c *fiber.Ctx) error
}

type playerHandler struct {
	logger        *logrus.Logger
	playerService services.PlayerService
}

func NewPlayerHandler(logger *logrus.Logger, playerService services.PlayerService) PlayerHandler {
	return &playerHandler{logger: logger, playerService: playerService}
}

func (h *playerHandler) Create(c *fiber.Ctx) error {
	var request api.PlayerCreateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		h.logger.Error("Error retrieving avatar file: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Avatar file is required",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		h.logger.Error("Error opening avatar file: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process avatar file",
		})
	}
	defer file.Close()

	avatar := api.Avatar{
		File:     file,
		Size:     fileHeader.Size,
		Filename: fileHeader.Filename,
	}

	request.Avatar = avatar

	err = h.playerService.Create(request)
	if err != nil {
		h.logger.Error("Error creating player: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create player",
		})
	}

	return c.Status(fiber.StatusCreated).JSON("Player created successfully")
}
