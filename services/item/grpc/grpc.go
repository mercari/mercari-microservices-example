package grpc

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"

	pkgconnect "github.com/mercari/mercari-microservices-example/pkg/connect"
	dbconnect "github.com/mercari/mercari-microservices-example/platform/db/proto/protoconnect"
	"github.com/mercari/mercari-microservices-example/services/item/proto/protoconnect"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	dbClient := dbconnect.NewDBServiceClient(
		pkgconnect.NewInsecureClient(),
		"http://db.db.svc.cluster.local:5000",
	)

	svc := &server{
		dbClient: dbClient,
	}
	return pkgconnect.RunServer(ctx, port, logger, func(opts ...connect.HandlerOption) (string, http.Handler) {
		return protoconnect.NewItemServiceHandler(svc, opts...)
	})
}
