package main

import (
	"bloc/database"
	"bloc/routes"
	"bloc/storage"
	"bloc/utils/config"
	"bloc/utils/env"
	"bloc/utils/tokens"
	"embed"
	"flag"
	"net/http"
	"os"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

//go:embed view/dist/*
var views embed.FS

func init() {
	// Configure logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fileConf := flag.String("config", "", "Path to the config file")
	flag.Parse()

	// Init configuration
	config.Init(env.Get("CONFIG_FILE", *fileConf))

	// Connect to database
	database.Connect()

	// Init oauth client
	if config.Conf.Oauth.Server != "" {
		cwauth.InitOauth2(oauth2.Config{
			ClientID:     config.Conf.Oauth.Id,
			ClientSecret: config.Conf.Oauth.Secret,
			RedirectURL:  config.Conf.Oauth.Callback,
		}, config.Conf.Oauth.Server)
	}

	// Init storage backend
	storage.Init(config.Conf.Storage.Driver)

	// Generate JWT token
	tokens.Init(env.Get("JWT_KEY", ""))
}

func main() {
	// Create fiber instance
	app := fiber.New(fiber.Config{
		BodyLimit: (1024 * 1024 * 1024) * 8, // Limit file upload size (8Gb)
	})

	// Include cors
	app.Use(cors.New())

	if env.Get("DEV_FRONT_URL", "") == "" {
		// Load view as static website
		app.Get("/", filesystem.New(filesystem.Config{
			Root:       http.FS(views),
			PathPrefix: "views/dist/app",
			Browse:     true,
		}))
	} else {
		app.Get("/*", func(c *fiber.Ctx) error {
			url := env.Get("DEV_FRONT_URL", "") + c.Params("*")
			err := proxy.Do(c, url)
			if err != nil {
				return err
			}

			return nil
		})
	}

	// Setup API routes
	routes.Api(app)

	log.Info().Err(app.Listen(config.Conf.Server.Address + ":" + config.Conf.Server.Port))
}
