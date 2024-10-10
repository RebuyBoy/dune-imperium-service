package app

import (
	"context"
	"dune-imperium-service/internal/config"
	"dune-imperium-service/internal/handlers"
	"dune-imperium-service/internal/repositories"
	"dune-imperium-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router      *fiber.App
	MongoClient *mongo.Client
	Logger      *logrus.Logger
}

func (app *App) Initialize() {
	app.Logger = config.SetupLogger()
	app.Logger.Info("Initializing the application...")

	app.Router = fiber.New()

	app.MongoClient = config.ConnectDB()
}

func (app *App) SetupRoutes() {
	gameRepo := repositories.NewGameRepository(app.MongoClient)
	gameService := services.NewGameService(app.Logger, gameRepo)
	gameController := handlers.NewResultHandler(app.Logger, gameService)

	api := app.Router.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/results")

	users.Get("/all", gameController.GetAll)
}

func (app *App) Run(address string) {
	defer func() {
		if err := app.MongoClient.Disconnect(context.Background()); err != nil {
			app.Logger.Error(err)
		}
	}()

	app.Logger.Infof("Starting server on %s", address)
	if err := app.Router.Listen(address); err != nil {
		app.Logger.Fatal(err)
	}
}
