// file: main.go
// version: 1.0.1
// guid: f1a2b3c4-d5e6-7f8a-9b0c-1d2e3f4a5b6c
// last-edited: 2026-07-19

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/falkcorp/gcommon/v2/internal/config"
	"github.com/falkcorp/gcommon/v2/internal/logging"
	"github.com/falkcorp/gcommon/v2/internal/metrics"
)

func main() {
	// Load configuration
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Initialize logger
	loggerConfig := &logging.Config{
		Level:      cfg.Logging.Level,
		Format:     cfg.Logging.Format,
		Output:     cfg.Logging.Output,
		TimeFormat: cfg.Logging.TimeFormat,
	}
	logger, err := logging.NewLogger(loggerConfig)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Initialize metrics
	metricsConfig := &metrics.Config{
		Enabled:   cfg.Metrics.Enabled,
		Namespace: cfg.Metrics.Namespace,
	}
	commonMetrics := metrics.NewCommonMetrics(metricsConfig)

	ctx := context.Background()
	logger.Info(ctx, "Starting gcommon server", "version", "1.0.0")

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Enable reflection for development
	if cfg.Server.EnableReflection {
		reflection.Register(grpcServer)
	}

	// Start gRPC server
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.GRPCPort))
	if err != nil {
		logger.Error(ctx, "Failed to listen", "error", err)
		os.Exit(1)
	}

	// Start server in goroutine
	go func() {
		logger.Info(ctx, "gRPC server starting",
			"host", cfg.Server.Host,
			"port", cfg.Server.GRPCPort,
		)

		commonMetrics.SetActiveConnections(1)

		if err := grpcServer.Serve(listen); err != nil {
			logger.Error(ctx, "gRPC server failed", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logger.Info(ctx, "Shutting down server...")
	grpcServer.GracefulStop()
	commonMetrics.SetActiveConnections(0)
	logger.Info(ctx, "Server shutdown complete")
}
