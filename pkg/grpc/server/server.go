package server

import (
	"fmt"
	"net"

	"github.com/rmscoal/Authenticator-API/pkg/logger"
	pb "github.com/rmscoal/Authenticator-API/pkg/pb"
	"google.golang.org/grpc"
)

var _defaultServerAddress = "50051"

type Server struct {
	pb.UnimplementedEmailerServer
	engine *grpc.Server
	port   string
	error  chan error
	logger logger.Interface
}

func New(l logger.Interface, opts ...Option) (*Server, error) {
	s := &Server{
		port:   _defaultServerAddress,
		logger: l,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s, nil
}

func (s *Server) start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		s.logger.Fatal(err, "gRPC failed to listen")
	}
	grpc_server := grpc.NewServer()
	pb.RegisterEmailerServer(grpc_server, s)
	s.logger.Info("gRPC server listening at: %v", lis.Addr().String())
	if err := grpc_server.Serve(lis); err != nil {
		s.logger.Fatal(err, "gRPC failed to server")
	}
	s.engine = grpc_server
}

func (s *Server) Notify() <-chan error {
	return s.error
}

func (s *Server) Shutdown() error {
	select {
	case <-s.error:
		return nil
	default:
	}

	s.engine.GracefulStop()

	return nil
}
