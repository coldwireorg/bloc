package files

import (
	"bloc/models"
	"bloc/storage"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
)

func Upload(c *fiber.Ctx) error {
	fileMultipart, err := c.FormFile("file") // We are getting the file sent via a form
	if err != nil {
		return errors.Handle(c, errors.ErrRequest, err)
	}

	// Get parent folder id
	parent := c.FormValue("parent")

	// Get encrypted key to store it, so the user never lost it
	// TODO: Add a regex to verify that a key is specified
	key := c.FormValue("key")

	if key == "" || parent == "" {
		return errors.Handle(c, errors.ErrRequest)
	}

	token, err := tokens.Parse(c.Cookies("token")) // Parse user's JWT token
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	p := models.Folder{Id: parent}
	p, err = p.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	// if the parent dir is not owned by the user, cancel upload
	if p.Owner != token.Username {
		return errors.Handle(c, errors.ErrPermission)
	}

	// TODO: check user's quota before upload

	var file = models.File{
		Id:         shortuuid.New(),
		Name:       fileMultipart.Filename,
		Size:       int(fileMultipart.Size),
		Type:       fileMultipart.Header["Content-Type"][0],
		IsFavorite: false,
		Key:        key,
		Parent:     parent,
		Owner:      token.Username,
	}

	// Upload file using the configured storage driver
	err = storage.Driver.Create(file.Id, fileMultipart)
	if err != nil {
		return errors.Handle(c, errors.ErrUnknown, err)
	}

	err = file.Create()
	if err != nil {
		storage.Driver.Delete(file.Id)
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	return errors.Handle(c, errors.Success, file)
}
