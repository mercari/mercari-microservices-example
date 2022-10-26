package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/mercari/mercari-microservices-example/services/authority/proto/protoconnect"
	customerconnect "github.com/mercari/mercari-microservices-example/services/customer/proto/protoconnect"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	customerClient := customerconnect.NewCustomerServiceClient(
		newInsecureClient(),
		"http://customer.customer.svc.cluster.local:5000",
	)

	svc := &server{
		customerClient: customerClient,
		logger:         logger.WithName("server"),
	}

	mux := http.NewServeMux()
	path, handler := protoconnect.NewAuthorityServiceHandler(svc)
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

func newInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				// If you're also using this client for non-h2c traffic, you may want
				// to delegate to tls.Dial if the network isn't TCP or the addr isn't
				// in an allowlist.
				return net.Dial(network, addr)
			},
			ReadIdleTimeout:  10 * time.Second,
			PingTimeout:      10 * time.Second,
			WriteByteTimeout: 10 * time.Second,
		},
	}
}
