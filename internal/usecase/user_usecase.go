package usecase

import (
	"context"
	"fmt"
	"user-service/internal/entity"
)

type UserRepository interface {
	GetById(ctx context.Context, id int) (*entity.User, error)
}

type UserUsecase struct {
	repo UserRepository
}

func New(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo}
}

func (uc *UserUsecase) GetByID(ctx context.Context, id int) (*entity.User, error) {
	user, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UserService - Usecase - GetById - repo.GetByID: %w", err)
	}
	return user, nil
}
