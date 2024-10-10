package app

import (
	"context"
	"myapp/internal/config"
	"myapp/internal/controllers"
	"myapp/internal/repositories"
	"myapp/internal/services"

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
	// Initialize Logger
	app.Logger = config.SetupLogger()
	app.Logger.Info("Initializing the application...")

	// Initialize Fiber Router
	app.Router = fiber.New()

	// Connect to MongoDB
	app.MongoClient = config.ConnectDB()
}

func (app *App) SetupRoutes() {
	userRepo := repositories.NewUserRepository(app.MongoClient)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, app.Logger)

	api := app.Router.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")

	users.Get("/", userController.GetUsers)
	users.Get("/:id", userController.GetUserByID)
	users.Post("/", userController.CreateUser)
	users.Put("/:id", userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)
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
