package service

import (
	"github.com/kahuri1/message-processor/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) CreateMessage(m model.Message) error {

	id, err := s.repo.Create(m)
	if err != nil {
		log.Errorf("failed to create message: %w", err)
	}

	log.Infof("message with id %d created", id)

	return nil
}
