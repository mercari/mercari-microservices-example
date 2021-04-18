package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/gateway/proto"
)

func RunServer(ctx context.Context, port int, grpcPort int) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	conn, err := grpc.DialContext(ctx, fmt.Sprintf(":%d", grpcPort), opts...)
	if err != nil {
		return fmt.Errorf("failed to dial grpc server: %w", err)
	}

	if err := proto.RegisterGatewayServiceHandlerClient(ctx, mux, proto.NewGatewayServiceClient(conn)); err != nil {
		return fmt.Errorf("failed to create a grpc client: %w", err)
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("failed to serve http server: %w", err)
	case <-ctx.Done():
		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown http server: %w", err)
		}

		if err := <-errCh; err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to close http server: %w", err)
		}

		return nil
	}
}
