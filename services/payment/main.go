package main

import (
	"context"
	"fmt"
	"os/signal"

	"golang.org/x/sys/unix"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer stop()

	select {
	case <-ctx.Done():
		fmt.Println("shutting down...")
	}
}
