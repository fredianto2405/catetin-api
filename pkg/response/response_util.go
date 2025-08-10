package response

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

func Respond(c *gin.Context, statusCode int, success bool, message string, data interface{}, meta *Meta) {
	c.JSON(statusCode, Response{
		Success: success,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func NewMeta(total, page, limit int) *Meta {
	if limit <= 0 {
		limit = 10
	}

	if page <= 0 {
		page = 1
	}

	pageCount := (total + limit - 1) / limit

	return &Meta{
		Total:     total,
		Page:      page,
		PageCount: pageCount,
	}
}
