package main

import (
	"UserManagment/gen/go/protos/wearable/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"sync"
)

func main() {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", options...)
	if err != nil {
		log.Fatalf("failed to dial grpc server %v", err)
	}

	defer conn.Close()

	grpcClient := wearable.NewWearableServiceClient(conn)

	stream, err := grpcClient.BeatsPerMinute(context.Background(), &wearable.BeatsPerMinuteRequest{
		Uuid: "sample-uuid",
	})

	if err != nil {
		log.Fatalf("stream failed %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			val, err := stream.Recv()

			fmt.Println(val)

			if err == io.EOF {
				return
			}

			if err != nil {
				log.Fatalf("stream failed in goroutine %v", err)
			}

			fmt.Println(val)
		}
	}()

	wg.Wait()
}
