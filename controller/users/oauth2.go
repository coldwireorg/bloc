package users

import (
	"bloc/models"
	"bloc/utils"
	"bloc/utils/errs"
	"time"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Oauth2Callback(c *fiber.Ctx) error {
	code := c.Query("code")

	idToken, accessToken := cwauth.Callback(code)

	tokenIsValid := cwauth.CheckToken(idToken, accessToken)
	if !tokenIsValid {
		return c.JSON(errs.Internal)
	}

	claims := cwauth.GetClaims(idToken)

	var usr = models.User{
		Username:   claims.Username,
		PrivateKey: claims.PrivateKey,
		PublicKey:  claims.PublicKey,
	}

	exist := usr.Exist()
	if !exist {
		err := usr.Create()
		if err != nil {
			log.Err(err).Msg(err.Error())
			return c.JSON(errs.Internal)
		}
	}

	utils.SetCookie(c, "access_token", accessToken, time.Now().Add(time.Hour*6))
	utils.SetCookie(c, "token", idToken, time.Now().Add(time.Hour*6))

	return c.Redirect("/app")
}
