package grpc

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	pkggrpc "github.com/mercari/mercari-microservices-example/pkg/grpc"
	db "github.com/mercari/mercari-microservices-example/platform/db/proto"
	"github.com/mercari/mercari-microservices-example/services/customer/proto"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	conn, err := grpc.DialContext(ctx, "db.db.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial db server: %w", err)
	}

	dbClient := db.NewDBServiceClient(conn)
	svc := &server{
		dbClient: dbClient,
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterCustomerServiceServer(s, svc)
	}).Start(ctx)
}
