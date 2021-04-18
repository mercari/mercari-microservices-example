package grpc

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"

	pkggrpc "github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/catalog/proto"
	customer "github.com/mercari/go-conference-2021-spring-office-hour/services/customer/proto"
	item "github.com/mercari/go-conference-2021-spring-office-hour/services/item/proto"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	cconn, err := grpc.DialContext(ctx, "customer.customer.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial customer grpc server: %w", err)
	}

	iconn, err := grpc.DialContext(ctx, "item.item.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial item grpc server: %w", err)
	}

	svc := &server{
		customerClient: customer.NewCustomerServiceClient(cconn),
		itemClient:     item.NewItemServiceClient(iconn),
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterCatalogServiceServer(s, svc)
	}).Start(ctx)
}
