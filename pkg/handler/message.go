package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/message-processor/pkg/model"
	log "github.com/sirupsen/logrus"
)

// Создание нового сообщения
func (h *Handler) CreateMessage(c *gin.Context) {

	var message model.Message

	d, err := c.GetRawData()

	err = json.Unmarshal(d, &message)
	if err != nil {
		log.Errorf("unmarshal handlerError")

		return
	}

	err = h.service.CreateMessage(message)

	log.Info("message created")
	//if err := c.ShouldBindJSON(&message); err != nil {
	//
	//	return
	//}
}

// Получение сообщения по ID
func (h *Handler) GetMessage(c *gin.Context) {
	// Логика получения сообщения
}
