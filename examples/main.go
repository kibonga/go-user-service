package main

import (
	users "UserManagment/gen/go/protos/user/v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type myUserService struct {
	users.UnimplementedUserServiceServer
}

func (userService *myUserService) GetUser(_ context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	return &users.GetUserResponse{
		User: &users.User{
			Uuid:     req.Uuid,
			FullName: "Pavle Djurdjic",
		},
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	users.RegisterUserServiceServer(grpcServer, &myUserService{})
	grpcServer.Serve(listener)
}
