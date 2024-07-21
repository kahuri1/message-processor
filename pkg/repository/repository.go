package repository

import "github.com/jmoiron/sqlx"

type MessageHTTPAPI interface {
}

type KafkaAPI interface {
}

type DatabaseAPI interface {
}

type LoggingAPI interface {
}

type Repository struct {
	MessageHTTPAPI
	KafkaAPI
	DatabaseAPI
	LoggingAPI
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
