package service

import "github.com/kahuri1/message-processor/pkg/repository"

type MessageHTTPAPI interface {
}

type KafkaAPI interface {
}

type DatabaseAPI interface {
}

type LoggingAPI interface {
}

type Service struct {
	MessageHTTPAPI
	KafkaAPI
	DatabaseAPI
	LoggingAPI
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
