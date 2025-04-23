package uspsgo

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

var programCtx atomic.Value
var programStop atomic.Value

func init() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	ctx, cancel := context.WithCancel(context.Background())
	sigChannel := make(chan os.Signal, 3)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigChannel
		fmt.Printf("got signal: %s", s.String())
		cancel()
	}()
	programCtx.Store(ctx)
	programStop.Store(cancel)
}

func Context() context.Context {
	return programCtx.Load().(context.Context)
}

func CancelContext() {
	programStop.Load().(context.CancelFunc)()
}
