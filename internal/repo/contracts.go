package repo

import (
	"context"

	"github.com/keshvan/user-service-sstu-forum/internal/entity"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*entity.User, error)
}
