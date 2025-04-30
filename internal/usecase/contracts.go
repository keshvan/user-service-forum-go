package usecase

import (
	"context"
)

type UserUsecase interface {
	GetUsernamesByIds(ctx context.Context, ids []int64) (map[int64]string, error)
}
