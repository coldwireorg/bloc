package users

import (
	"bloc/models"
	"bloc/utils"
	"bloc/utils/errs"
	"time"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
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

	var root = models.Folder{
		Id:   shortuuid.New(),
		Name: "root",
	}

	var usr = models.User{
		Username:   claims.Username,
		PrivateKey: claims.PrivateKey,
		PublicKey:  claims.PublicKey,
	}

	exist := usr.Exist()
	if !exist {
		// Create root folder
		root.Create()

		// create user
		err := usr.Create()
		if err != nil {
			log.Err(err).Msg(err.Error())
			return c.JSON(errs.Internal)
		}

		usr.SetRoot(root.Id)        // Set root folder of the user
		root.SetOwner(usr.Username) // Set the owner of the root folder
	}

	utils.SetCookie(c, "access_token", accessToken, time.Now().Add(time.Hour*6))
	utils.SetCookie(c, "token", idToken, time.Now().Add(time.Hour*6))

	return c.Redirect("/app")
}
