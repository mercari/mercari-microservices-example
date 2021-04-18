package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/mercari/go-conference-2021-spring-office-hour/pkg/logger"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/gateway/grpc"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/gateway/http"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	l, err := logger.New()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create logger: %s", err)
		if ferr != nil {
			// Unhandleable, something went wrong...
			panic(fmt.Sprintf("failed to write log:`%s` original error is:`%s`", ferr, err))
		}
		return 1
	}
	glogger := l.WithName("gateway")

	grpcErrCh := make(chan error, 1)
	go func() {
		grpcErrCh <- grpc.RunServer(ctx, 5000, glogger.WithName("grpc"))
	}()

	httpErrCh := make(chan error, 1)
	go func() {
		httpErrCh <- http.RunServer(ctx, 4000, 5000)
	}()

	select {
	case err := <-grpcErrCh:
		glogger.Error(err, "failed to serve grpc server")
		return 1
	case err := <-httpErrCh:
		glogger.Error(err, "failed to serve http server")
		return 1
	case <-ctx.Done():
		glogger.Info("shutting down...")
		return 0
	}
}
