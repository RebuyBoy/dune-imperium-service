package user

import (
	"dune-imperium-service/internal/handlers"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupPlayerRoutes(router fiber.Router, logger *logrus.Logger, userService services.PlayerService) {
	userHandler := handlers.NewPlayerHandler(logger, userService)

	usersGroup := router.Group("/users")

	usersGroup.Get("/all", userHandler.GetAll)
}
