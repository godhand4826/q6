package serve

import "github.com/gin-gonic/gin"

func SendError(c *gin.Context, code int, message string) {
	e := HTTPError{
		Code:    code,
		Message: message,
	}
	c.JSON(code, e)
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
