package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	authoritypb "github.com/mercari/mercari-microservices-example/services/authority/proto"
	catalogpb "github.com/mercari/mercari-microservices-example/services/catalog/proto"
)

func RunServer(ctx context.Context, port int) error {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: false,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	authorityConn, err := grpc.DialContext(ctx, "authority.authority.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial to authority grpc server: %w", err)
	}
	if err = authoritypb.RegisterAuthorityServiceHandlerClient(ctx, mux, authoritypb.NewAuthorityServiceClient(authorityConn)); err != nil {
		return fmt.Errorf("failed to create a authority grpc client: %w", err)
	}

	catalogConn, err := grpc.DialContext(ctx, "catalog.catalog.svc.cluster.local:5000", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial to catalog grpc server: %w", err)
	}
	if err := catalogpb.RegisterCatalogServiceHandlerClient(ctx, mux, catalogpb.NewCatalogServiceClient(catalogConn)); err != nil {
		return fmt.Errorf("failed to create a catalog grpc client: %w", err)
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
