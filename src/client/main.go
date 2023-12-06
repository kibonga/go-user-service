package main

import (
	"UserManagment/gen/go/protos/wearable/v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	stream, err := client.ConsumeBeatsPerMinute(context.Background())
	if err != nil {
		log.Fatalf("failed to create stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		err = stream.Send(&wearable.ConsumeBeatsPerMinuteRequest{
			Uuid:   "abc",
			Value:  uint32(i * rand.Intn(100)),
			Minute: uint32(i + 1),
		})

		if err != nil {
			log.Fatalf("failed to send: %v", err)
		}

		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to close: %v", err)
	}

	log.Println(resp.GetTotal())
}
