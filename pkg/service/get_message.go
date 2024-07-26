package service

import (
	"github.com/kahuri1/message-processor/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) GetMessage(id int) (*model.GetMessage, error) {

	message, err := s.repo.GetMessage(id)
	if err != nil {
		log.Errorf("failed to get mess: %w", err)
	}

	log.Infof("message: %d ", message)

	return message, nil
}
