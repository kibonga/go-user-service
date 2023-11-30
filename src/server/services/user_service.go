package services

import (
	users "UserManagment/gen/go/protos/user/v1"
	"context"
)

type UserServiceImpl struct {
	users.UnimplementedUserServiceServer
}

func (userService *UserServiceImpl) GetUser(_ context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	return &users.GetUserResponse{
		User: &users.User{
			Uuid:     req.Uuid,
			FullName: "Pavle Djurdjic",
		},
	}, nil
}
