package grpc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	pkgconnect "github.com/mercari/mercari-microservices-example/pkg/connect"
	"github.com/mercari/mercari-microservices-example/pkg/connect/interceptor"
	"github.com/mercari/mercari-microservices-example/services/authority/proto/protoconnect"
	customerconnect "github.com/mercari/mercari-microservices-example/services/customer/proto/protoconnect"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	customerClient := customerconnect.NewCustomerServiceClient(
		pkgconnect.NewInsecureClient(),
		"http://customer.customer.svc.cluster.local:5000",
	)

	svc := &server{
		customerClient: customerClient,
		logger:         logger.WithName("server"),
	}

	interceptors := connect.WithInterceptors(interceptor.NewRequestLogger(logger))
	mux := http.NewServeMux()
	path, handler := protoconnect.NewAuthorityServiceHandler(svc, interceptors)
	mux.Handle(path, handler)
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	return server.ListenAndServe()
}
