package main

import (
	"context"
	"fmt"
	"go-gin-blog-api/config"
	"go-gin-blog-api/internal/server"
	"go-gin-blog-api/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	config.Load()
	if err := logger.Init(config.AppConfig.Env); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Log.Sync()
	r, err := server.New()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.AppConfig.Port),
		Handler: r,
	}

	go func() {
		logger.Log.Info("starting server",
			zap.String("addr", srv.Addr),
			zap.String("env", config.AppConfig.Env),
		)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("forced shutdown", zap.Error(err))
	}

	logger.Log.Info("server exited gracefully")
}
