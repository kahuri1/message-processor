package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Обработка ошибок
func (h *Handler) HandleError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
