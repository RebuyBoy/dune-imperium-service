package server

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"hs-fake-api/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App    *fiber.App
	Config *config.Config
	Logger *log.Logger
}

func NewServer(cfg *config.Config, logger *log.Logger) *Server {
	app := fiber.New(
		fiber.Config{
			AppName:           cfg.AppName,
			EnablePrintRoutes: true,
		})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	server := &Server{
		App:    app,
		Config: cfg,
		Logger: logger,
	}

	return server
}

func (s *Server) Start() error {
	address := ":" + s.Config.Port
	s.Logger.Printf("%s is running on port %s", s.Config.AppName, s.Config.Port)
	return s.App.Listen(address)
}
