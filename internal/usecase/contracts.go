package usecase

import (
	"context"

	"github.com/keshvan/user-service-sstu-forum/internal/entity"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id int) (*entity.User, error)
}
