package app

import (
	"context"
	"dune-imperium-service/internal/configs"
	"dune-imperium-service/internal/db"
	"dune-imperium-service/internal/server"
	"dune-imperium-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Server      *fiber.App
	MongoClient *mongo.Client
	Logger      *logrus.Logger
	Cfg         *configs.Config
}

func (app *App) Initialize() {
	app.Logger = configs.SetupLogger()

	err := godotenv.Load()
	if err != nil {
		app.Logger.Fatal("Error loading .env file")
	}

	app.Cfg, err = configs.GetConfig()
	if err != nil {
		app.Logger.Fatal("Error getting config: ", err)
	}

	app.MongoClient, err = db.ConnectMongoDB(app.Cfg.MongoURI)
	if err != nil {
		app.Logger.Fatal("Error connecting to MongoDB: ", err)
	}

	serviceContainer := services.NewServiceContainer(app.Logger, app.MongoClient)

	app.Server = server.NewServer(app.Logger, app.Cfg, serviceContainer)
}

func (app *App) Run() {
	defer func() {
		if err := app.MongoClient.Disconnect(context.Background()); err != nil {
			app.Logger.Error(err)
		}
	}()

	if err := app.Server.Listen(":" + app.Cfg.Port); err != nil {
		app.Logger.Fatal(err)
	}
}
