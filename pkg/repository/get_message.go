package repository

import (
	"database/sql"

	"github.com/kahuri1/message-processor/pkg/model"
)

func (r *Repository) GetMessage(id int) (*model.GetMessage, error) {
	var message model.GetMessage
	sqlQuery := `SELECT message_id, sender_id, content, created_at, status 
            FROM messages
            WHERE message_id = $1`
	err := r.db.QueryRow(sqlQuery, id).Scan(&message.ID, &message.SenderID, &message.Content, &message.CreatedAt, &message.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Сообщение не найдено
		}
		return nil, err // Ошибка при выполнении запроса
	}
	return &message, nil
}
