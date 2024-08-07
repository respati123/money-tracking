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

type PaginationModel struct {
	TotalData   int `json:"total_data"`
	TotalPage   int `json:"total_page"`
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
}

type PaginationResponse struct {
	Data     interface{}
	Metadata PaginationModel
}

type PaginationRequest struct {
	Page    int                    `json:"page"`
	PerPage int                    `json:"per_page"`
	Filter  map[string]interface{} `json:"filter"`
}
