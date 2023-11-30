package main

import (
	users "UserManagment/gen/go/protos/user/v1"
	"UserManagment/gen/go/protos/wearable/v1"
	"UserManagment/src/server/services"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	users.RegisterUserServiceServer(grpcServer, &services.UserServiceImpl{})
	wearable.RegisterWearableServiceServer(grpcServer, &services.WearableServiceImpl{})
	fmt.Println("Server started...")
	grpcServer.Serve(listener)
}
