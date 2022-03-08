package main

import (
	"bloc/config"
	"bloc/database"
	"bloc/routes"
	"bloc/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// Configure logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Init configuration
	config.Init()

	// Connect to database
	go database.Connect()
}

func main() {
	isReady := utils.Migrate()

	// Create fiber instance
	app := fiber.New(fiber.Config{
		BodyLimit: (1024 * 1024 * 1024) * 8, // Limit file upload size (8Gb)
	})

	// Include cors
	app.Use(cors.New())

	// Setup API routes
	routes.Api(app)

	// Wait for database to be clean/ready
	if isReady {
		log.Info().Err(app.Listen(config.Conf.Server.Address + ":" + config.Conf.Server.Port))
	}
}
