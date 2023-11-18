package grpc

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/nikolaevv/airtraffic/internal/action/bookings"
	"github.com/nikolaevv/airtraffic/internal/action/flights"
	"github.com/nikolaevv/airtraffic/internal/adaptor"
	"github.com/nikolaevv/airtraffic/internal/service/grpc/pb"
	"github.com/nikolaevv/airtraffic/pkg/converter"

	"github.com/pkg/errors"
)

func Init(cont *adaptor.Container) pb.AirTrafficServiceServer {
	return &Service{cont: cont}
}

type Service struct {
	pb.UnimplementedAirTrafficServiceServer
	cont *adaptor.Container
}

func (s Service) GetFlights(ctx context.Context, _ *pb.GetFlightsRq) (*pb.GetFlightsRs, error) {
	act := flights.NewGetList(s.cont.GetFlightRepository())

	flightsList, err := act.Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get flights list")
	}

	res := &pb.GetFlightsRs{}
	err = copier.CopyWithOption(res, &flightsList, converter.DefaultConverterOptions)
	if err != nil {
		return nil, errors.Wrap(err, "copy flights list")
	}

	return res, nil
}

func (s Service) GetBooking(ctx context.Context, req *pb.GetBookingRq) (*pb.Booking, error) {
	act := bookings.NewGet(s.cont.GetBookingRepository())

	booking, err := act.Do(ctx, int(req.Id))
	if err != nil {
		return nil, errors.Wrap(err, "get booking by id")
	}

	res := &pb.Booking{}
	err = copier.CopyWithOption(res, &booking, converter.DefaultConverterOptions)
	if err != nil {
		return nil, errors.Wrap(err, "copy booking")
	}

	return res, err
}
