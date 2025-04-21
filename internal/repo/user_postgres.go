package repo

import (
	"context"
	"fmt"

	"github.com/keshvan/go-common-forum/postgres"
	"github.com/keshvan/user-service-sstu-forum/internal/entity"
)

type userRepository struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *userRepository {
	return &userRepository{pg}
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	row := r.pg.Pool.QueryRow(ctx, "SELECT id, username, is_admin, password_hash FROM users WHERE id = $1", id)

	var u entity.User

	if err := row.Scan(&u.ID, &u.Username, &u.IsAdmin, &u.PasswordHash); err != nil {
		return nil, fmt.Errorf("UserRepository - GetByID - row.Scan(): %w", err)
	}

	return &u, nil
}
