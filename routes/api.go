package routes

import (
	"bloc/controller/files"
	"bloc/controller/files/shares"
	"bloc/controller/users"
	"bloc/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	api := app.Group("/api")

	// Users > /api/user
	user := api.Group("/user")

	// Auth > /api/user/auth
	auth := user.Group("/auth")
	auth.Post("/register", users.Register)
	auth.Post("/login", users.Login)
	auth.Post("/logout", users.Logout)

	// Oauth2 > /api/user/auth/oauth2
	oauth := auth.Group("/oauth2")
	oauth.Get("/callback", users.Oauth2Callback)

	// files > /api/file
	file := api.Group("/file", middlewares.Auth)
	file.Post("/", files.Upload)          // Upload file
	file.Delete("/", files.Delete)        // Delete file
	file.Get("/", files.List)             // List files
	file.Get("/:id", files.Download)      // Download file
	file.Put("/favorite", files.Favorite) // set favorite

	// Shares  > /api/file/share
	share := file.Group("/share")
	share.Post("/", shares.Add)      // Share a file
	share.Delete("/", shares.Revoke) // Revoke a share
	share.Get("/", shares.List)      // List share
}
