package routes

import (
	"dune-imperium-service/internal/routes/content"
	"dune-imperium-service/internal/routes/result"
	"dune-imperium-service/internal/routes/user"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger, services *services.Container) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	result.SetupResultRoutes(v1, logger, services.ResultService)
	user.SetupPlayerRoutes(v1, logger, services.PlayerService)
	content.SetupContentRoutes(v1, logger, services.ContentService)

}
