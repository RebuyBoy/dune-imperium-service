package handlers

import (
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ResultHandler struct {
	logger      *logrus.Logger
	gameService services.GameService
}

func NewResultHandler(logger *logrus.Logger, gameService services.GameService) *ResultHandler {
	return &ResultHandler{logger: logger, gameService: gameService}
}
func (h *ResultHandler) GetAll(c *fiber.Ctx) error {
	h.logger.Info("Fetching all results")
	results := h.gameService.GetAll()
	return c.JSON(results)
}
