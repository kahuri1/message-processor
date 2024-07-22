package handler

import (
	"github.com/gin-gonic/gin"
)

// Создание нового сообщения
func (h *Handler) CreateMessage(c *gin.Context) {

	if err := c.ShouldBindJSON(&message); err != nil {

		return
	}
}

// Получение сообщения по ID
func (h *Handler) GetMessage(c *gin.Context) {
	// Логика получения сообщения
}
