package grpc

import (
	"context"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"

	pkggrpc "github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/category/proto"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterCategoryServiceServer(s, &server{})
	}).Start(ctx)
}
