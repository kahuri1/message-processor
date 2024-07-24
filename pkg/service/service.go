package service

import (
	"github.com/kahuri1/message-processor/pkg/model"
	log "github.com/sirupsen/logrus"
)

type kafka interface {
}

type repo interface {
	Create(message model.Message) (int64, error)
	CreateUser(message model.User) (int64, error)

}

type Service struct {
	kafka kafka
	repo  repo
}

func NewService(repo repo, kafka kafka) *Service {
	log.Info("service init")

	return &Service{
		repo:  repo,
		kafka: kafka,
	}
}
