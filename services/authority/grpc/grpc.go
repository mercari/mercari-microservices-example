package grpc

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"

	pkgconnect "github.com/mercari/mercari-microservices-example/pkg/connect"
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

	return pkgconnect.RunServer(ctx, port, logger, func(opts ...connect.HandlerOption) (string, http.Handler) {
		return protoconnect.NewAuthorityServiceHandler(svc, opts...)
	})
}
