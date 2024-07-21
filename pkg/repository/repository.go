package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
