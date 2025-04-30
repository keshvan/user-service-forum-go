package usecase

import (
	"context"
	"fmt"

	"github.com/keshvan/user-service-sstu-forum/internal/repo"
)

type userUsecase struct {
	repo repo.UserRepository
}

func New(repo repo.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) GetUsernamesByIds(ctx context.Context, ids []int64) (map[int64]string, error) {
	if len(ids) == 0 {
		return make(map[int64]string), nil
	}

	usernames, err := u.repo.GetUsernamesByIds(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("UserService - UserUsecase - GetUsernamesByIds - repo.GetUsernamesByIds: %w", err)
	}

	return usernames, nil
}
