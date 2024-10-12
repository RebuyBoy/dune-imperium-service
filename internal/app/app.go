package app

import (
	"context"
	"dune-imperium-service/internal/configs"
	"dune-imperium-service/internal/db"
	"dune-imperium-service/internal/server"
	"dune-imperium-service/internal/services"
	"dune-imperium-service/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Logger      *logrus.Logger
	HttpServer  *fiber.App
	MongoClient *mongo.Client
	MinioClient *minio.Client
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

	app.MongoClient, err = db.MongoDBClient(app.Cfg.MongoURI)
	if err != nil {
		app.Logger.Fatal("Error connecting to MongoDB: ", err)
	}

	app.MinioClient, err = storage.MinioClient(app.Cfg.Minio)
	if err != nil {
		app.Logger.Fatal("Error initializing MinIO client: ", err)
	}

	deps := services.ServiceDependencies{
		Logger:      app.Logger,
		MongoClient: app.MongoClient,
		MinioClient: app.MinioClient,
	}

	serviceContainer := services.NewServiceContainer(deps)

	app.HttpServer = server.NewServer(app.Logger, app.Cfg, serviceContainer)
}

func (app *App) Run() {
	defer func() {
		if err := app.MongoClient.Disconnect(context.Background()); err != nil {
			app.Logger.Error(err)
		}
	}()

	if err := app.HttpServer.Listen(":" + app.Cfg.Port); err != nil {
		app.Logger.Fatal(err)
	}
}
