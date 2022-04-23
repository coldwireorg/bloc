package errors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Status  int    `json:"-"`
	Code    string `json:"status"`
	Message string `json:"message"`

	Content interface{} `json:"content"`
	Error   string      `json:"error"`
}

func Handle(c *fiber.Ctx, args ...interface{}) error {
	var res Response

	if args[0] == nil {
		log.Panic().Msg("The response must be specified !!")
	}

	res = Response{
		Status:  args[0].(Response).Status,
		Code:    args[0].(Response).Code,
		Message: args[0].(Response).Message,
	}

	if len(args) > 1 {
		switch args[1].(type) {
		case error:
			log.Err(args[1].(error)).Msg(args[1].(error).Error()) // Print error
			res.Error = args[1].(error).Error()
		default:
			res.Content = args[1]
		}
	}

	return c.Status(res.Status).JSON(res)
}

var (
	/* INTERNAL */
	Success = Response{
		Status:  200,
		Code:    "SUCCESS",
		Message: "Request treated with success!",
	}
	ErrRequest = Response{
		Status:  400,
		Code:    "ERROR_REQUEST",
		Message: "The sever can't handle the request",
	}
	ErrBody = Response{
		Status:  400,
		Code:    "ERROR_BODY",
		Message: "The server can't parse body",
	}
	ErrUnknown = Response{
		Status:  500,
		Code:    "ERROR_UNKNOWN",
		Message: "An unknown error occured",
	}
	ErrPermission = Response{
		Status:  500,
		Code:    "ERROR_PERMISSION",
		Message: "You can't access this content",
	}

	/* DATABASE */
	ErrDatabaseCreate = Response{
		Status:  500,
		Code:    "ERROR_DB",
		Message: "Failed to put data in the database",
	}
	ErrDatabaseUpdate = Response{
		Status:  500,
		Code:    "ERROR_DB",
		Message: "Failed to update data in the database",
	}
	ErrDatabaseRemove = Response{
		Status:  500,
		Code:    "ERROR_DB",
		Message: "Failed to remove data from the database",
	}
	ErrDatabaseNotFound = Response{
		Status:  504,
		Code:    "ERROR_DB",
		Message: "Failed to find data in the database",
	}

	/* AUTH */
	ErrAuth = Response{
		Status:  500,
		Code:    "ERROR_AUTH",
		Message: "An issue occured during authentication",
	}
	ErrAuthExist = Response{
		Status:  403,
		Code:    "ERROR_AUTH_EXIST",
		Message: "This user already exist",
	}
	ErrAuthPassword = Response{
		Status:  403,
		Code:    "ERROR_AUTH_PASSWORD",
		Message: "Wrong password",
	}
)
