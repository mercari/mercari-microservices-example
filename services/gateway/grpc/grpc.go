package grpc

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"

	pkggrpc "github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc"
	authority "github.com/mercari/go-conference-2021-spring-office-hour/services/authority/proto"
	catalog "github.com/mercari/go-conference-2021-spring-office-hour/services/catalog/proto"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/gateway/proto"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	aconn, err := grpc.DialContext(ctx, "authority.authority.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial authority grpc server: %w", err)
	}

	cconn, err := grpc.DialContext(ctx, "catalog.catalog.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial catalog grpc server: %w", err)
	}

	svc := &server{
		authorityClient: authority.NewAuthorityServiceClient(aconn),
		catalogClient:   catalog.NewCatalogServiceClient(cconn),
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterGatewayServiceServer(s, svc)
	}).Start(ctx)
}
