package base

type BaseResponse struct {
	Status       int         `json:"status"`
	ErrorMessage *string     `json:"error_message"`
	Data         interface{} `json:"data"`
}
