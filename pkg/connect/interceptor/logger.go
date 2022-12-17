package interceptor

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"
	connectcontext "github.com/mercari/mercari-microservices-example/pkg/connect/context"
	"google.golang.org/grpc/status"
)

func NewRequestLogger(logger logr.Logger) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			ctx = connectcontext.SetRequestID(ctx)
			reqid := connectcontext.GetRequestID(ctx)

			logger.WithValues(
				"method", req.Spec().Procedure,
				"request_id", reqid,
			).Info("grpc request")

			res, err := next(ctx, req)

			logger.WithValues(
				"method", req.Spec().Procedure,
				"code", status.Code(err),
				"request_id", reqid,
			).Info("finished")

			return res, err
		})
	})
}
