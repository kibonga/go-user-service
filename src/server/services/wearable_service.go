package services

import (
	wearable "UserManagment/gen/go/protos/wearable/v1"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"math/rand"
	"time"
)

type WearableServiceServerImpl struct {
	wearable.UnimplementedWearableServiceServer
}

func (wearableServer *WearableServiceServerImpl) BeatsPerMinute(
	req *wearable.BeatsPerMinuteRequest,
	stream wearable.WearableService_BeatsPerMinuteServer) error {

	// On the server side, server receives the request and then returns the stream response
	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended")
		default:
			time.Sleep(1 * time.Second)
			value := 30 + rand.Int31n(80)

			err := stream.SendMsg(&wearable.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			})
			if err != nil {
				return status.Error(codes.Canceled, "Stream has ended")
			}
		}
	}

}

func (wearableServer *WearableServiceServerImpl) ConsumeBeatsPerMinute(stream wearable.WearableService_ConsumeBeatsPerMinuteServer) error {
	var total uint32

	// On server side, server receives a value from the stream (which is sent by the client)
	for {
		val, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wearable.ConsumeBeatsPerMinuteResponse{
				Total: total,
			})
		}

		if err != nil {
			return err
		}

		fmt.Println(val.GetUuid(), val.GetMinute(), val.GetValue())
		total++
	}
}

func (wearableServer *WearableServiceServerImpl) CalculateBeatsPerMinute(stream wearable.WearableService_CalculateBeatsPerMinuteServer) error {
	var count, total uint32

	// On server side, server receives values from the stream, computes them and returns a stream response
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		total += req.GetValue()
		fmt.Println("received: ", req.GetValue())

		count++
		if (count % 5) == 0 {
			fmt.Println("total:", total, "sending: ", float32(total)/5)
			if err := stream.Send(&wearable.CalculateBeatsPerMinuteResponse{
				Average: float32(total) / 5,
			}); err != nil {
				return nil
			}

			total = 0
		}
	}
}
