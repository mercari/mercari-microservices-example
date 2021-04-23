package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const metadataAuthorizationKey = "authorization"

func NewAuthTokenPropagator() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			authmd := md.Get(metadataAuthorizationKey)
			if len(authmd) > 0 {
				ctx = metadata.AppendToOutgoingContext(ctx, metadataAuthorizationKey, authmd[0])
			}
		}

		return handler(ctx, req)
	}
}
