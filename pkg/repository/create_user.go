package repository

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/kahuri1/message-processor/pkg/model"
)

func (r *Repository) CreateUser(user model.User) (int64, error) {
	sql, args, err := squirrel.
		Insert("users").
		Columns("username", "email", "created_at").
		Values(user.Username, user.Email, squirrel.Expr("NOW()")).
		Suffix("RETURNING user_id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to build user create query: %w", err)
	}

	var id int64

	err = r.db.QueryRow(sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to process user create query: %w", err)
	}

	return id, nil
}
