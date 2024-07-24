package handler

import (
	"net/http"

	"github.com/kahuri1/message-processor/pkg/model"

	"github.com/gin-gonic/gin"
)

type messageService interface {
	CreateMessage(m model.Message) error
	CreateUser(m model.User) error
}

type Handler struct {
	service messageService
}

func Newhandler(service messageService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	messagesGroup := router.Group("/messages")
	{
		messagesGroup.POST("/", h.CreateMessage) // Создание сообщения
		messagesGroup.GET("/:id", h.GetMessage)  // Получение сообщения по ID
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/", h.createUser)
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
