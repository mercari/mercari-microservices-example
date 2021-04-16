package grpc

import (
	"context"

	"google.golang.org/grpc"

	pkggrpc "github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/balance/proto"
)

func Run(ctx context.Context, port int) error {
	return pkggrpc.NewServer(port, func(s *grpc.Server) {
		proto.RegisterBalanceServiceServer(s, nil)
	}).Start(ctx)
}
