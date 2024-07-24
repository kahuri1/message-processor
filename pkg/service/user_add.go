package service

import (
	"github.com/kahuri1/message-processor/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) CreateUser(m model.User) error {

	id, err := s.repo.CreateUser(m)
	if err != nil {
		log.Errorf("failed to create user: %w", err)
	}

	log.Infof("message with id %d user", id)

	return nil
}
