package grpc

import (
	"context"

	"github.com/nikolaevv/airtraffic/internal/action/flights"
	"github.com/nikolaevv/airtraffic/internal/adaptor"
	"github.com/nikolaevv/airtraffic/internal/service/grpc/pb"

	"github.com/pkg/errors"
)

func Init(cont *adaptor.Container) pb.AirTrafficServiceServer {
	return &Service{cont: cont}
}

type Service struct {
	pb.UnimplementedAirTrafficServiceServer
	cont *adaptor.Container
}

func (s Service) GetFlights(ctx context.Context, req *pb.GetFlightsRq) (*pb.GetFlightsRs, error) {
	act := flights.NewGetList(s.cont.GetFlightRepository())

	flights, err := act.Do(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := &pb.GetFlightsRs{}
	for _, flight := range flights {
		res.Flights = append(res.Flights, &pb.Flight{
			Id:                 int64(flight.ID),
			ScheduledDeparture: flight.ScheduledDeparture.Format("2006-01-02 15:04:05"),
			ScheduledArrival:   flight.ScheduledArrival.Format("2006-01-02 15:04:05"),
			Status:             string(flight.Status),
			AircraftId:         int64(flight.AircraftID),
			ActualDeparture:    flight.ActualDeparture.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}
