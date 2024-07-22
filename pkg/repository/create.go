package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/message-processor/pkg/model"
)

func (r *Repository) Create(message model.Message) (int64, error) {
	sql, args, err := sq.
		Insert("messages").
		Columns("sender_id", "recipient_id", "content").
		Values(message.SenderID, message.RecipientID, message.Content).
		Suffix("RETURNING message_id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to build news create query: %w", err)
	}

	var id int64

	err = r.db.QueryRow(sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to proccess news create query: %w", err)
	}

	return id, nil
}
