package services

import (
	wearable "UserManagment/gen/go/protos/wearable/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"time"
)

type WearableServiceImpl struct {
	wearable.UnimplementedWearableServiceServer
}

func (wearableServer *WearableServiceImpl) BeatsPerMinute(
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
