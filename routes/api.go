package routes

import (
	"bloc/controller/auth"
	"bloc/controller/files"
	"bloc/controller/folders"
	"bloc/controller/shares"
	"bloc/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	api := app.Group("/api")

	// Auth > /api/user/auth
	authPoint := api.Group("/auth")
	authPoint.Post("/register", middlewares.Allowed, auth.Register)
	authPoint.Post("/login", middlewares.Allowed, auth.Login)
	authPoint.All("/logout", auth.Logout)

	// Oauth2 > /api/user/auth/oauth2
	oauth := authPoint.Group("/oauth2")
	oauth.Get("/", auth.Oauth2)
	oauth.Get("/callback", auth.Oauth2Callback)

	// files > /api/file
	file := api.Group("/file", middlewares.Auth)
	file.Post("/", files.Upload)              // Upload file
	file.Delete("/:id", files.Delete)         // Delete file
	file.Get("/list/:folder", files.List)     // List files in a folder
	file.Get("/download/:id", files.Download) // Download file
	file.Put("/favorite/:id", files.Favorite) // set favorite

	// folders > /api/folder
	folder := api.Group("/folder", middlewares.Auth)
	folder.Post("/", folders.Create)
	folder.Put("/:id", folders.Move)
	folder.Delete("/:id", folders.Delete) // Delete file

	// Shares  > /api/share
	share := api.Group("/share", middlewares.Auth)
	share.Post("/", shares.Add)         // Share a file
	share.Delete("/:id", shares.Revoke) // Revoke a share
	share.Get("/", shares.List)         // List share
}
