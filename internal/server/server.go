// internal/server/server.go

package server

import (
	"dune-imperium-service/internal/configs"
	"dune-imperium-service/internal/routes"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewServer(logger *logrus.Logger, cfg *configs.Config, svc *services.Container) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:           cfg.AppName,
		EnablePrintRoutes: true,
	})

	routes.SetupRoutes(app, logger, svc)

	return app
}
