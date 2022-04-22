package routes

import (
	"bloc/controller/auth"
	"bloc/controller/files"
	"bloc/controller/folders"
	"bloc/controller/shares"
	"bloc/controller/user"
	"bloc/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	api := app.Group("/api")

	// Auth > /api/auth
	authRoute := api.Group("/auth")
	authRoute.Post("/register", middlewares.IsOauthEnabled, auth.Register)
	authRoute.Post("/login", middlewares.IsOauthEnabled, auth.Login)
	authRoute.All("/logout", auth.Logout)

	// Oauth2 > /api/auth/oauth2
	oauthRoute := authRoute.Group("/oauth2")
	oauthRoute.Get("/", auth.Oauth2)
	oauthRoute.Get("/callback", auth.Oauth2Callback)

	// User > /api/user/
	userRoute := api.Group("/user", middlewares.IsAuthenticated)
	userRoute.Get("/info", user.Info)

	// files > /api/file
	fileRoute := api.Group("/file", middlewares.IsAuthenticated)
	fileRoute.Post("/", files.Upload)              // Upload file
	fileRoute.Put("/:id", files.Move)              // Move file
	fileRoute.Delete("/:id", files.Delete)         // Delete file
	fileRoute.Get("/list/:folder", files.List)     // List files in a folder
	fileRoute.Get("/download/:id", files.Download) // Download file
	fileRoute.Put("/favorite/:id", files.Favorite) // set favorite

	// folders > /api/folder
	folderRoute := api.Group("/folder", middlewares.IsAuthenticated)
	folderRoute.Post("/", folders.Create)
	folderRoute.Put("/:id", folders.Move)
	folderRoute.Delete("/:id", folders.Delete) // Delete file

	// Shares  > /api/share
	shareRoute := api.Group("/share", middlewares.IsAuthenticated)
	shareRoute.Post("/", shares.Add)         // Share a file
	shareRoute.Delete("/:id", shares.Revoke) // Revoke a share
	shareRoute.Get("/", shares.List)         // List share
}
