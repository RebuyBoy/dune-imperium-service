package handlers

import (
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ContentHandler struct {
	logger         *logrus.Logger
	contentService *services.ContentService
}

func NewContentHandler(
	logger *logrus.Logger,
	contentService *services.ContentService,
) *ContentHandler {
	return &ContentHandler{
		logger:         logger,
		contentService: contentService,
	}
}

func (h *ContentHandler) Get(c *fiber.Ctx) error {
	h.logger.Info("GetContent")
	return c.SendString("GetContent")
}

func (h *ContentHandler) Create(c *fiber.Ctx) error {
	h.logger.Info("GetContent")
	return c.SendString("CreateContent")
}
