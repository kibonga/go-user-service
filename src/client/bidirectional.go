package main

import (
	"UserManagment/gen/go/protos/wearable/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	"time"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := wearable.NewWearableServiceClient(conn)
	stream, err := client.CalculateBeatsPerMinute(context.Background())
	if err != nil {
		log.Fatalf("failed to establish connection")
	}

	for i := 0; i < 10; i++ {
		err = stream.Send(&wearable.CalculateBeatsPerMinuteRequest{
			Uuid:   "abc",
			Minute: uint32(i * rand.Intn(5)),
			Value:  uint32(i * rand.Intn(100)),
		})

		if err != nil {
			log.Fatalf("failed to send: %v", err)
		}

		time.Sleep(100 * time.Millisecond)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("failed to close: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to receive from stream: %v", err)
		}

		fmt.Println("average:", resp.GetAverage())
	}
}
