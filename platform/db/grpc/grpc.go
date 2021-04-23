package grpc

import (
	"context"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"

	pkggrpc "github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc"
	"github.com/mercari/go-conference-2021-spring-office-hour/platform/db/db"
	"github.com/mercari/go-conference-2021-spring-office-hour/platform/db/proto"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	svc := &server{
		db: db.New(),
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterDBServiceServer(s, svc)
	}).Start(ctx)
}
