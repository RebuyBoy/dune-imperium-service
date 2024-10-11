package handlers

import (
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PlayerHandler interface {
	GetAll(c *fiber.Ctx) error
}

type playerHandler struct {
	logger        *logrus.Logger
	playerService services.PlayerService
}

func NewPlayerHandler(logger *logrus.Logger, playerService services.PlayerService) PlayerHandler {
	return &playerHandler{logger: logger, playerService: playerService}
}
func (h *playerHandler) GetAll(c *fiber.Ctx) error {
	h.logger.Info("Fetching all players")
	players := h.playerService.GetAll()
	return c.JSON(players)
}
