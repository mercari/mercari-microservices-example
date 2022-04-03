package grpc

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	pkggrpc "github.com/mercari/mercari-microservices-example/pkg/grpc"
	"github.com/mercari/mercari-microservices-example/services/authority/proto"
	customer "github.com/mercari/mercari-microservices-example/services/customer/proto"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	conn, err := grpc.DialContext(ctx, "customer.customer.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial grpc server: %w", err)
	}

	customerClient := customer.NewCustomerServiceClient(conn)

	svc := &server{
		customerClient: customerClient,
		logger:         logger.WithName("server"),
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterAuthorityServiceServer(s, svc)
	}).Start(ctx)
}
