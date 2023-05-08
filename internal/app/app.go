package app

import (
	"context"
	"log"

	"user-microservice/internal/api/handler"
	"user-microservice/internal/api/router"
	"user-microservice/internal/config"
	"user-microservice/internal/repository"
	"user-microservice/internal/service"
	"user-microservice/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func Run() error {
	config.InitEnvConfigs()

	pool, err := database.NewPostgresDB(
		context.Background(),
		database.Config{
			Host:     config.EnvConfig.DBHost,
			Port:     config.EnvConfig.DBPort,
			Username: config.EnvConfig.DBUsername,
			Name:     config.EnvConfig.DBName,
			Password: config.EnvConfig.DBPassword,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	app := fiber.New()

	repos := repository.NewRepository(pool)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := router.NewServer(app, handlers)

	server.Router()
	app.Listen(":" + config.EnvConfig.LocalServerPort)

	return nil
}
