package repo

import (
	"context"
	"fmt"

	"github.com/keshvan/go-common-forum/postgres"
)

type userRepository struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) UserRepository {
	return &userRepository{pg}
}

func (r *userRepository) GetUsernamesByIds(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		return make(map[int64]string), nil
	}

	rows, err := r.pg.Pool.Query(ctx, "SELECT id, username FROM users WHERE id = ANY($1)", ids)
	if err != nil {
		return nil, fmt.Errorf("UserRepository - GetUsernamesByIds - r.pg.Pool.Query: %w", err)
	}
	defer rows.Close()

	usernames := make(map[int64]string, len(ids))
	for rows.Next() {
		var id int64
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			return nil, fmt.Errorf("UserRepository - GetUsernamesByIds - rows.Scan: %w", err)
		}
		usernames[id] = username
	}

	return usernames, nil
}
