package handlers

import (
	"dune-imperium-service/internal/dto/api"
	"dune-imperium-service/internal/mappers"
	"dune-imperium-service/internal/models"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PlayerHandler struct {
	logger        *logrus.Logger
	playerService *services.PlayerService
}

func NewPlayerHandler(logger *logrus.Logger, playerService *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{logger: logger, playerService: playerService}
}

func (h *PlayerHandler) GetById(c *fiber.Ctx) error {
	player, err := h.playerService.GetById(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}

	response := mappers.ToPlayerResponse(player)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *PlayerHandler) Create(c *fiber.Ctx) error {
	request, err := h.parseCreateRequest(c)
	if err != nil {
		return h.handleError(c, fiber.StatusBadRequest, "Invalid request", err)
	}

	avatar, err := h.processAvatar(c)
	if err != nil {
		return h.handleError(c, fiber.StatusBadRequest, "Failed to process avatar", err)
	}
	request.Avatar = avatar

	newPlayer, err := h.playerService.Create(c.Context(), request)
	if err != nil {
		if err.Error() == "nickname already exists" {
			return h.handleError(c, fiber.StatusConflict, "Nickname already exists", err)
		}
		return h.handleError(c, fiber.StatusInternalServerError, "Failed to create player", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"player": fiber.Map{
			"id":        newPlayer.ID,
			"nickname":  newPlayer.Nickname,
			"avatarURL": newPlayer.AvatarURL,
		},
	})
}

func (h *PlayerHandler) GetNames(c *fiber.Ctx) error {
	names, err := h.playerService.GetNames(c.Context())
	if err != nil {
		return h.handleError(c, fiber.StatusInternalServerError, "Failed to get player names", err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"names": names})
}

func (h *PlayerHandler) parseCreateRequest(c *fiber.Ctx) (api.PlayerCreateRequest, error) {
	var request api.PlayerCreateRequest
	if err := c.BodyParser(&request); err != nil {
		return request, err
	}
	return request, nil
}

func (h *PlayerHandler) processAvatar(c *fiber.Ctx) (*models.FileData, error) {
	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &models.FileData{
		Filename: fileHeader.Filename,
		Size:     fileHeader.Size,
		Content:  file,
	}, nil
}

func (h *PlayerHandler) handleError(c *fiber.Ctx, status int, message string, err error) error {
	h.logger.Error(message+": ", err)
	return c.Status(status).JSON(fiber.Map{"error": message})
}
