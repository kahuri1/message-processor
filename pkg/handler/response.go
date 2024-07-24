package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handlerError struct {
	message string `json: "message"`
}

// Обработка ошибок
func newErorResnonse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, handlerError{message})
}
