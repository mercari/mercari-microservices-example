package grpc

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"

	pkggrpc "github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc"
	db "github.com/mercari/go-conference-2021-spring-office-hour/platform/db/proto"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/item/proto"
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
		proto.RegisterItemServiceServer(s, svc)
	}).Start(ctx)
}
