package interceptor

import (
	"context"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func NewRequestLogger(logger logr.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.WithValues(
			"method", info.FullMethod,
		).Info("grpc request")

		res, err := handler(ctx, req)

		logger.WithValues(
			"method", info.FullMethod,
			"code", status.Code(err),
		).Info("finished")

		return res, err
	}
}
