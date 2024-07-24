package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/kahuri1/message-processor/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) createUser(c *gin.Context) {

	var user model.User

	d, err := c.GetRawData()
	if err != nil {
		log.Info("user created error: ", err)
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Errorf("unmarshal handlerError")

		return
	}

	err = h.service.CreateUser(user)

	log.Info("user created: ", err)

}
