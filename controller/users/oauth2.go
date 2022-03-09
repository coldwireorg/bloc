package users

import (
	"bloc/utils"
	"time"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
)

func Oauth2Callback(c *fiber.Ctx) error {
	code := c.Query("code")

	idToken, accessToken := cwauth.Callback(code)

	utils.SetCookie(c, "access_token", accessToken, time.Now().Add(time.Hour*6))
	utils.SetCookie(c, "token", idToken, time.Now().Add(time.Hour*6))

	return c.Redirect("/app")
}
