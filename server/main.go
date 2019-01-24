package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/push/server/config"

	"github.com/Noah-Huppert/golog"
	"github.com/gorilla/mux"
)

func main() {
	// {{{1 Setup logger
	logger := golog.NewStdLogger("push-server")

	// {{{1 Load configuration
	cfg, err := config.NewConfigFromEnv()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// {{{1 Make handle ctrl+c
	ctx, ctxCancel := context.WithCancel(context.Background())

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	go func() {
		<-interruptChan
		ctxCancel()
	}()

	// {{{1 Setup HTTP server
	// {{{2 Setup handlers
	handler := mux.NewRouter()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler: handler,
	}

	// {{{2 Start server
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			logger.Fatalf("error serving API: %s", err.Error())
		}
		logger.Debug("stopped API server")
	}()

	// {{{2 Handle context close
	go func() {
		<-ctx.Done()
		logger.Debug("stopping API server")
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Fatalf("error shutting down API server: %s",
				err.Error())
		}
	}()

	logger.Debugf("started API server on port %d", cfg.HTTPPort)
}
