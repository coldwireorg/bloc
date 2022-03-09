package utils

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx, name string, value string, expire time.Time) {
	uri := strings.Split(c.Hostname(), ":")

	c.Cookie(&fiber.Cookie{
		Name:        name,
		Value:       value,
		Expires:     expire,
		Domain:      uri[0],
		HTTPOnly:    true,
		SessionOnly: false,
		Secure:      c.Secure(),
	})
}

func DelCookie(c *fiber.Ctx, name string) {
	SetCookie(c, name, "", time.Now().Add(time.Second*-30))
}
