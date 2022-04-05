package files

import (
	"bloc/models"
	"bloc/storage"
	"bloc/utils"
	"bloc/utils/errs"
	"bloc/utils/tokens"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
)

func Upload(c *fiber.Ctx) error {
	fileMultipart, err := c.FormFile("file") // We are getting the file sent via a form
	if err != nil {
		// If we can't get the sent file, we sen a 500 error
		return c.Status(500).JSON(errs.BadRequest)
	}

	// Get parent folder
	parent := c.FormValue("parent")

	// Get encrypted key to store it, so the user never lost it
	// TODO: Add a regex to verify that a key is specified
	key := c.FormValue("key")

	if key == "" || parent == "" {
		return c.Status(500).JSON(errs.BadRequest)
	}

	token, err := tokens.Parse(c.Cookies("token")) // Parse user's JWT token
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(errs.BadRequest)
	}

	// TODO: check user's quota before upload

	var file = models.File{
		Id:         shortuuid.New(),
		Name:       fileMultipart.Filename,
		Size:       int(fileMultipart.Size),
		IsFavorite: false,
		Key:        key,
		Parent:     parent,
		Owner:      token.Username,
	}

	// Upload file using the configured storage driver
	err = storage.Driver.Create(file.Id, fileMultipart)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(errs.Internal)
	}

	err = file.Create()
	if err != nil {
		log.Println(err)
		storage.Driver.Delete(file.Id)
		return c.Status(500).JSON(errs.Internal)
	}

	return c.JSON(utils.Reponse{
		Success: true,
		Data:    file,
	})
}
