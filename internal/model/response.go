package model

type Response struct {
	ResponseMessage string      `json:"response_message"`
	ResponseData    interface{} `json:"response_data"`
	ResponseError   string      `json:"response_error"`
}

type ResponseInterface struct {
	Message string
	Data    interface{}
}
