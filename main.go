package main

import (
	"bloc/database"
	"bloc/routes"
	"bloc/storage"
	"bloc/utils"
	"bloc/utils/config"
	"bloc/utils/env"
	"flag"
	"os"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

func init() {
	// Configure logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fileConf := flag.String("config", "", "Path to the config file")
	flag.Parse()

	// Init configuration
	config.Init(env.Get("CONFIG_FILE", *fileConf))

	// Connect to database
	go database.Connect()

	// Init oauth client
	if config.Conf.Oauth.Server != "" {
		go cwauth.InitOauth2(oauth2.Config{
			ClientID:     config.Conf.Oauth.Id,
			ClientSecret: config.Conf.Oauth.Secret,
			RedirectURL:  config.Conf.Oauth.Callback,
		}, config.Conf.Oauth.Server)
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
