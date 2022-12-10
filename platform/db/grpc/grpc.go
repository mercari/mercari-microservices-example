package grpc

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"

	pkgconnect "github.com/mercari/mercari-microservices-example/pkg/connect"
	"github.com/mercari/mercari-microservices-example/platform/db/db"
	"github.com/mercari/mercari-microservices-example/platform/db/proto/protoconnect"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	svc := &server{
		db: db.New(),
	}

	return pkgconnect.RunServer(ctx, port, logger, func(opts ...connect.HandlerOption) (string, http.Handler) {
		return protoconnect.NewDBServiceHandler(svc, opts...)
	})
}
