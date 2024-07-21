package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Проверка работоспособности сервиса
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
