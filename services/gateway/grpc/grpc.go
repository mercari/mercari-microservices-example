package grpc

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"

	pkggrpc "github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc"
	authority "github.com/mercari/go-conference-2021-spring-office-hour/services/authority/proto"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/gateway/proto"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	conn, err := grpc.DialContext(ctx, "authority.authority.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial grpc server: %w", err)
	}

	authorityClient := authority.NewAuthorityServiceClient(conn)

	svc := &server{authorityClient: authorityClient}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterGatewayServiceServer(s, svc)
	}).Start(ctx)
}
