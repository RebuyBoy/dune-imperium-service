package handlers

import (
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler interface {
	GetAll(c *fiber.Ctx) error
}

type userHandler struct {
	logger      *logrus.Logger
	gameService services.UserService
}

func NewUserHandler(logger *logrus.Logger, gameService services.UserService) UserHandler {
	return &userHandler{logger: logger, gameService: gameService}
}
func (h *userHandler) GetAll(c *fiber.Ctx) error {
	h.logger.Info("Fetching all results")
	results := h.gameService.GetAll()
	return c.JSON(results)
}
