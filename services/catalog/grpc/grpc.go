package grpc

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"

	pkgconnect "github.com/mercari/mercari-microservices-example/pkg/connect"
	"github.com/mercari/mercari-microservices-example/services/catalog/proto/protoconnect"
	customerconnect "github.com/mercari/mercari-microservices-example/services/customer/proto/protoconnect"
	itemconnect "github.com/mercari/mercari-microservices-example/services/item/proto/protoconnect"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	itemclient := itemconnect.NewItemServiceClient(
		pkgconnect.NewInsecureClient(),
		"http://item.item.svc.cluster.local:5000",
	)

	customerclient := customerconnect.NewCustomerServiceClient(
		pkgconnect.NewInsecureClient(),
		"http://customer.customer.svc.cluster.local:5000",
	)

	svc := &server{
		customerClient: customerclient,
		itemClient:     itemclient,
	}
	return pkgconnect.RunServer(ctx, port, logger, func(opts ...connect.HandlerOption) (string, http.Handler) {
		return protoconnect.NewCatalogServiceHandler(svc, opts...)
	})
}
