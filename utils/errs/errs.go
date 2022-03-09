package errs

import (
	"bloc/utils"
)

var (
	Internal = utils.Reponse{
		Success: false,
		Error:   "Internal server error",
	}
	BadRequest = utils.Reponse{
		Success: false,
		Error:   "Bad request, please fill all needed fields",
	}
	AuthBadPassword = utils.Reponse{
		Success: false,
		Error:   "Bad password",
	}
	AuthNoKeypair = utils.Reponse{
		Success: false,
		Error:   "No keys provided",
	}
	AuthNameAlreadyTaken = utils.Reponse{
		Success: false,
		Error:   "Username already taken",
	}
)
