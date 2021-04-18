package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/authority/grpc"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpc.Run(ctx, 5000)
	}()

	select {
	case err := <-errCh:
		fmt.Println(err.Error())
		return 1
	case <-ctx.Done():
		fmt.Println("shutting down...")
		return 0
	}
}
