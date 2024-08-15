package model

type Response struct {
	ResponseCode    int         `json:"response_code"`
	ResponseMessage string      `json:"response_message"`
	ResponseData    interface{} `json:"response_data,omitempty"`
	ResponseError   string      `json:"response_error,omitempty"`
}

type PaginationRequest struct {
	Page    int                    `json:"page"`
	PerPage int                    `json:"per_page"`
	Filter  map[string]interface{} `json:"filter"`
}

type PaginationMetadata struct {
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
}

type PaginationResponse struct {
	Data     interface{}        `json:"data"`
	Metadata PaginationMetadata `json:"metadata"`
}

type ResponseInterface struct {
	StatusCode int
	Error      error
	Message    string
	Data       interface{}
}
