package interceptor

import (
	"context"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	grpccontext "github.com/mercari/mercari-microservices-example/pkg/grpc/context"
)

func NewRequestLogger(logger logr.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = grpccontext.SetRequestID(ctx)
		reqid := grpccontext.GetRequestID(ctx)

		logger.WithValues(
			"method", info.FullMethod,
			"request_id", reqid,
		).Info("grpc request")

		res, err := handler(ctx, req)

		logger.WithValues(
			"method", info.FullMethod,
			"code", status.Code(err),
			"request_id", reqid,
		).Info("finished")

		return res, err
	}
}
