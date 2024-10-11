package handlers

import (
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ResultHandler struct {
	logger        *logrus.Logger
	resultService services.ResultService
}

func NewResultHandler(logger *logrus.Logger, resultService services.ResultService) *ResultHandler {
	return &ResultHandler{logger: logger, resultService: resultService}
}

func (h *ResultHandler) GetAll(c *fiber.Ctx) error {
	h.logger.Info("Fetching all results")
	results := h.resultService.GetAll()
	return c.JSON(results)
}
