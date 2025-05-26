package response

import (
	"backend/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	Data        any    `json:"data,omitempty"`
	Credentials any    `json:"credentials,omitempty"`
}

func RespondWithError(c *gin.Context, code int, message string) {
	logger.Error("[ERROR]", message)
	c.AbortWithStatusJSON(code, Response{
		Success: false,
		Message: message,
	})
}

func RespondWithSuccess(c *gin.Context, code int, message string, data any, credentials any) {
	logger.Error("[SUCCESS]", message)
	c.AbortWithStatusJSON(code, Response{
		Success:     true,
		Message:     message,
		Data:        data,
		Credentials: credentials,
	})
}

func OK(c *gin.Context, data any) {
	RespondWithSuccess(c, http.StatusOK, "OK", data, nil)
}

func BadRequest(c *gin.Context, err error) {
	RespondWithError(c, http.StatusBadRequest, err.Error())
}
