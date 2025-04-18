package repo

import (
	"context"

	"github.com/keshvan/go-common-forum/postgres"
	"github.com/keshvan/user-service-forum-go/internal/entity"
)

type UserRepository struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

func (r *UserRepository) GetById(ctx context.Context, id int) (*entity.User, error) {
	sql := r.pg.Pool.QueryRow(ctx, "SELECT id, username, is_admin, password_hash FROM users WHERE id = $1", id)

	var u entity.User
	err := sql.Scan(&u.ID, &u.Username, &u.IsAdmin, &u.PasswordHash)

	if err != nil {
		return nil, err
	}
	return &u, nil
}
