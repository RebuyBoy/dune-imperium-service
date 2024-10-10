package routes

import (
	"dune-imperium-service/internal/routes/results"
	"dune-imperium-service/internal/routes/users"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger, services *services.Container) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	results.SetupResultsRoutes(v1, logger, services.GameService)
	users.SetupUsersRoutes(v1, logger, services.UserService)

}
