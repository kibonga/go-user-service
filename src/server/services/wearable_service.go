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
