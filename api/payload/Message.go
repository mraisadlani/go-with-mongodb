package payload

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func ResponseJSON(status int, message string, data interface{}) Response {
	res := Response{
		Status: status,
		Message: message,
		Data: data,
	}

	return res
}

func ResponseError(status int, message string) ErrorResponse {
	res := ErrorResponse{
		Status: status,
		Message: message,
	}

	return res
}