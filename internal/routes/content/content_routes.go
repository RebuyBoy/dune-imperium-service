package content

import (
	"dune-imperium-service/internal/handlers"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupContentRoutes(router fiber.Router, logger *logrus.Logger, contentService *services.ContentService) {
	contentHandler := handlers.NewContentHandler(logger, contentService)

	contentGroup := router.Group("/content")

	contentGroup.Get("/:id", contentHandler.GetById)
	contentGroup.Get("/type/:type", contentHandler.GetByType)
	contentGroup.Post("/", contentHandler.Create)
}
