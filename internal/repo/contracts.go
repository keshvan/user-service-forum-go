package repo

import (
	"context"
)

type UserRepository interface {
	GetUsernamesByIds(ctx context.Context, ids []int64) (map[int64]string, error)
}
