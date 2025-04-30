package grpc

import (
	"context"
	"fmt"

	userpb "github.com/keshvan/protos-forum/user"
	"github.com/keshvan/user-service-sstu-forum/internal/usecase"
	"google.golang.org/grpc"
)

type serverAPI struct {
	userpb.UnimplementedUserServiceServer
	usecase usecase.UserUsecase
}

func Register(grpcServer *grpc.Server, usecase usecase.UserUsecase) {
	userpb.RegisterUserServiceServer(grpcServer, &serverAPI{usecase: usecase})
}

func (s *serverAPI) GetUsernames(ctx context.Context, req *userpb.GetUsernamesRequest) (*userpb.GetUsernamesResponse, error) {
	usernames, err := s.usecase.GetUsernamesByIds(ctx, req.UserIds)
	if err != nil {
		return nil, fmt.Errorf("UserService - UserServer - GetUsernames - usecase.GetUsernamesByIds: %w", err)
	}

	return &userpb.GetUsernamesResponse{Usernames: usernames}, nil
}
