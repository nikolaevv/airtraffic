package grpc

import (
	"context"

	"github.com/nikolaevv/airtraffic/internal/action/boarding_pass"
	"github.com/nikolaevv/airtraffic/internal/action/bookings"
	"github.com/nikolaevv/airtraffic/internal/action/flights"
	"github.com/nikolaevv/airtraffic/internal/adaptor"
	"github.com/nikolaevv/airtraffic/internal/model"
	"github.com/nikolaevv/airtraffic/internal/service/grpc/pb"
	"github.com/nikolaevv/airtraffic/pkg/converter"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Init(cont *adaptor.Container) pb.AirTrafficServiceServer {
	return &Service{cont: cont}
}

type Service struct {
	pb.UnimplementedAirTrafficServiceServer
	cont *adaptor.Container
}

func (s Service) GetFlights(ctx context.Context, _ *pb.GetFlightsRq) (*pb.GetFlightsRs, error) {
	l := s.cont.GetLogger()
	act := flights.NewGetList(s.cont.GetRepository())

	flightsList, err := act.Do(ctx)
	if err != nil {
		l.Error("get flights list", zap.Error(err))
		return nil, errors.Wrap(err, "get flights list")
	}

	res := &pb.GetFlightsRs{}
	err = copier.CopyWithOption(&res.Flights, &flightsList, converter.DefaultConverterOptions)
	if err != nil {
		l.Error("copy flights list", zap.Error(err))
		return nil, errors.Wrap(err, "copy flights list")
	}

	return res, nil
}

func (s Service) GetBooking(ctx context.Context, req *pb.GetBookingRq) (*pb.Booking, error) {
	l := s.cont.GetLogger()
	act := bookings.NewGet(s.cont.GetRepository())

	booking, err := act.Do(ctx, int(req.Id))
	if err != nil {
		l.Error("get booking by id", zap.Error(err))
		return nil, errors.Wrap(err, "get booking by id")
	}

	res := &pb.Booking{}
	err = copier.CopyWithOption(res, &booking, converter.DefaultConverterOptions)
	if err != nil {
		l.Error("copy booking", zap.Error(err))
		return nil, errors.Wrap(err, "copy booking")
	}

	return res, err
}

func (s Service) CreateBoardingPass(ctx context.Context, req *pb.CreateBoardingPassRq) (*pb.BoardingPass, error) {
	l := s.cont.GetLogger()
	act := boarding_pass.NewCreate(s.cont.GetRepository())

	boardingPass, err := act.Do(ctx, int(req.TicketFlightId), int(req.SeatId))
	if err != nil {
		l.Error("create boarding pass", zap.Error(err))
		return nil, errors.Wrap(err, "create boarding pass")
	}

	res := &pb.BoardingPass{}
	err = copier.CopyWithOption(res, &boardingPass, converter.DefaultConverterOptions)
	if err != nil {
		l.Error("copy boarding pass", zap.Error(err))
		return nil, errors.Wrap(err, "copy boarding pass")
	}

	return res, nil
}

func (s Service) BookTickets(ctx context.Context, req *pb.BookTicketsRq) (*pb.Booking, error) {
	l := s.cont.GetLogger()
	passengers := make([]model.Passenger, 0, len(req.Passengers))

	err := copier.Copy(&passengers, &req.Passengers)
	if err != nil {
		l.Error("copy passengers", zap.Error(err))
		return nil, errors.Wrap(err, "copy passengers")
	}

	act := bookings.NewCreate(s.cont.GetRepository())

	booking, err := act.Do(ctx, int(req.FlightId), passengers)
	if err != nil {
		l.Error("create booking", zap.Error(err))
		return nil, errors.Wrap(err, "create booking")
	}

	res := &pb.Booking{}
	err = copier.CopyWithOption(res, &booking, converter.DefaultConverterOptions)
	if err != nil {
		l.Error("copy booking", zap.Error(err))
		return nil, errors.Wrap(err, "copy booking")
	}

	return res, nil
}
