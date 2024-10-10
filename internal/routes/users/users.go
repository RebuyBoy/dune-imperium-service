package users

import (
	"dune-imperium-service/internal/handlers"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupUsersRoutes(router fiber.Router, logger *logrus.Logger, userService services.UserService) {
	userHandler := handlers.NewUserHandler(logger, userService)

	usersGroup := router.Group("/users")

	usersGroup.Get("/all", userHandler.GetAll)
}
