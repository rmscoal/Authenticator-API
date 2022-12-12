package server

import (
	"log"
	"net"

	"github.com/rmscoal/Authenticator-API/pkg/logger"
	pb "github.com/rmscoal/Authenticator-API/pkg/pb"
	"google.golang.org/grpc"
)

var _defaultServerAddress = "50051"

type Server struct {
	pb.UnimplementedEmailerServer
	port   string
	logger logger.Interface
}

func New(l logger.Interface, opts ...Option) (*Server, error) {
	s := &Server{
		port:   _defaultServerAddress,
		logger: l,
	}

	s.start()

	return s, nil
}

func (s *Server) start() {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpc_server := grpc.NewServer()
	pb.RegisterEmailerServer(grpc_server, s)
	log.Printf("Server listening at: %v", lis.Addr().String())
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
