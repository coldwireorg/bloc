package main

import (
	"bloc/config"
	"bloc/database"
	"bloc/routes"
	"bloc/storage"
	"bloc/utils"
	"os"

	"codeberg.org/coldwire/cwauth"
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

	// Init oauth client
	if config.Conf.Oauth.Server != "" {
		go cwauth.InitOauth2(config.Conf.Oauth.Config, config.Conf.Oauth.Server)
	}

	// Init storage backend
	storage.Init()
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
