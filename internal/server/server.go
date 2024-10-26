package server

import (
	"dune-imperium-service/internal/configs"
	"dune-imperium-service/internal/routes"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

func NewServer(
	logger *logrus.Logger,
	cfg *configs.Config,
	svc *services.Container,
) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:           cfg.AppName,
		EnablePrintRoutes: true,
	})
	app.Use(cors.New())

	routes.SetupRoutes(app, logger, svc)

	return app
}
