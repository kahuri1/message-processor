package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahuri1/message-processor/pkg/service"
)

type Handler struct {
	services *service.Service
}

func Newhandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	messagesGroup := router.Group("/messages")
	{
		messagesGroup.POST("/", h.CreateMessage) // Создание сообщения
		messagesGroup.GET("/:id", h.GetMessage)  // Получение сообщения по ID
	}

	statsGroup := router.Group("/stats")
	{
		statsGroup.GET("/", h.GetStats) // Получение статистики
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}
