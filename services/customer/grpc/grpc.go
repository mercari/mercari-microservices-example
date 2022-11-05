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
	dbconnect "github.com/mercari/mercari-microservices-example/platform/db/proto/protoconnect"
	"github.com/mercari/mercari-microservices-example/services/customer/proto/protoconnect"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	dbClient := dbconnect.NewDBServiceClient(
		pkgconnect.NewInsecureClient(),
		"http://db.db.svc.cluster.local:5000",
	)

	svc := &server{
		dbClient: dbClient,
	}

	interceptors := connect.WithInterceptors(interceptor.NewRequestLogger(logger))
	mux := http.NewServeMux()
	path, handler := protoconnect.NewCustomerServiceHandler(svc, interceptors)
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
