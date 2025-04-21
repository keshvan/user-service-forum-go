package usecase

import (
	"context"
	"fmt"

	"github.com/keshvan/user-service-sstu-forum/internal/entity"
	"github.com/keshvan/user-service-sstu-forum/internal/repo"
)

type userUsecase struct {
	repo repo.UserRepository
}

func New(repo repo.UserRepository) *userUsecase {
	return &userUsecase{repo}
}

func (uc *userUsecase) GetByID(ctx context.Context, id int) (*entity.User, error) {
	user, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UserService - Usecase - GetById - repo.GetByID: %w", err)
	}
	return user, nil
}
