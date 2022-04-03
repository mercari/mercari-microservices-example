package grpc

import (
	"context"

	"github.com/go-logr/logr"
	pkggrpc "github.com/mercari/mercari-microservices-example/pkg/grpc"
	"github.com/mercari/mercari-microservices-example/platform/db/db"
	"github.com/mercari/mercari-microservices-example/platform/db/proto"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	svc := &server{
		db: db.New(),
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterDBServiceServer(s, svc)
	}).Start(ctx)
}
