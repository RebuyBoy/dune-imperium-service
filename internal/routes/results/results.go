package results

import (
	"dune-imperium-service/internal/handlers"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupResultsRoutes(router fiber.Router, logger *logrus.Logger, resultService services.UserService) {
	resultHandler := handlers.NewResultHandler(logger, resultService)

	resultsGroup := router.Group("/results")

	resultsGroup.Get("/all", resultHandler.GetAll)
}
