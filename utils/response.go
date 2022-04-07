package utils

type Reponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}
