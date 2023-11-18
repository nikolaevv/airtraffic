package app

import (
	"fmt"
	"net"

	"github.com/nikolaevv/airtraffic/internal/adaptor"
	grpcService "github.com/nikolaevv/airtraffic/internal/service/grpc"
	"github.com/nikolaevv/airtraffic/internal/service/grpc/pb"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Server struct {
	cont *adaptor.Container
}

func (s *Server) Start() error {
	listener, err := net.Listen(
		s.cont.GetConfig().GetString("app.network"),
		fmt.Sprintf("%s:%s", s.cont.GetConfig().GetString("app.host"), s.cont.GetConfig().GetString("app.port")),
	)

	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterAirTrafficServiceServer(grpcServer, grpcService.Init(s.cont))
	err = grpcServer.Serve(listener)

	return errors.Wrap(err, "start grpc server")
}
