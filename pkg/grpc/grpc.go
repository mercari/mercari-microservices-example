package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-logr/logr"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"

	"github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc/interceptor"
)

type Server struct {
	server *grpc.Server
	port   int
}

func NewServer(port int, logger logr.Logger, register func(server *grpc.Server)) *Server {
	interceptors := []grpc.UnaryServerInterceptor{
		interceptor.NewRequestLogger(logger.WithName("request")),
	}

	opts := []grpc.ServerOption{
		middleware.WithUnaryServerChain(interceptors...),
	}

	server := grpc.NewServer(opts...)

	register(server)

	reflection.Register(server)
	channelz.RegisterChannelzServiceToServer(server)

	return &Server{
		server: server,
		port:   port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", s.port, err)
	}

	errCh := make(chan error, 1)
	go func() {
		if err := s.server.Serve(listener); err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("server has stopped with error: %w", err)
		}
		return nil

	case <-ctx.Done():
		s.server.GracefulStop()
		return <-errCh
	}
}
