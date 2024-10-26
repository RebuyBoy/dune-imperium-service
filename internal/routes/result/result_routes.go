package result

import (
	"dune-imperium-service/internal/handlers"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupResultRoutes(router fiber.Router, logger *logrus.Logger, resultService *services.ResultService) {
	resultHandler := handlers.NewResultHandler(logger, resultService)

	resultsGroup := router.Group("/results")

	resultsGroup.Get("/", resultHandler.Get)
	resultsGroup.Post("/", resultHandler.Save)
}
