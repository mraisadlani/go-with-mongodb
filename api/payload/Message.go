package payload

import "github.com/gin-gonic/gin"

type ResponseJSON struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"result"`
}

type ErrorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

type ResponsePagination struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"result"`
	Count int `json:"count"`
	Total int64 `json:"total"`
	Limit int64 `json:"limit"`
	Page int64 `json:"page"`
}

func Response(c *gin.Context, status int, message string, data interface{}) {
	res := ResponseJSON{
		Status: status,
		Message: message,
		Data: data,
	}

	c.JSON(status, res)
	return
}

func ResponsePage(c *gin.Context, status int, message string, data interface{}, count int, total int64, limit int64, page int64) {
	res := ResponsePagination{
		Status: status,
		Message: message,
		Data: data,
		Count: count,
		Total: total,
		Limit: limit,
		Page: page,
	}

	c.JSON(status, res)
	return
}

func ResponseError(c *gin.Context, status int, message string) {
	res := ErrorResponse{
		Status: status,
		Message: message,
	}

	c.JSON(status, res)
	return
}