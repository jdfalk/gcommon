// file: services/health/service.go
// version: 1.1.0
// guid: a1b2c3d4-e5f6-7a8b-9c0d-1e2f3a4b5c6d

package health

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"google.golang.org/grpc"

	commonpb "github.com/falkcorp/gcommon/pkg/commonpb"
	healthpb "github.com/falkcorp/gcommon/pkg/healthpb"
)

// Internal domain types for business logic
type InternalHealthStatus int

const (
	InternalHealthUnknown InternalHealthStatus = iota
	InternalHealthHealthy
	InternalHealthUnhealthy
	InternalHealthDegraded
)

type InternalHealthResult struct {
	CheckID   string
	Status    InternalHealthStatus
	CheckedAt time.Time
	Duration  time.Duration
	Message   string
	Error     error
	Details   map[string]string
}

// Service implements the health service with comprehensive health checking
type Service struct {
	healthpb.UnimplementedHealthServiceServer

	mu sync.RWMutex

	// Service status tracking (internal domain types)
	status    InternalHealthStatus // Use internal enum, not protobuf enum
	startTime time.Time            // Native Go time
	lastCheck time.Time            // Native Go time

	// Dependencies for health checking
	dependencies map[string]HealthChecker

	// Configuration
	config *Config
}

// convertToProtobufStatus converts internal status to protobuf status
func convertToProtobufStatus(internal InternalHealthStatus) healthpb.HealthStatus {
	switch internal {
	case InternalHealthHealthy:
		return healthpb.HealthStatus_HEALTH_STATUS_HEALTHY
	case InternalHealthUnhealthy:
		return healthpb.HealthStatus_HEALTH_STATUS_UNHEALTHY
	case InternalHealthDegraded:
		return healthpb.HealthStatus_HEALTH_STATUS_DEGRADED
	default:
		return healthpb.HealthStatus_HEALTH_STATUS_UNSPECIFIED
	}
}

// HealthChecker defines the interface for checking dependency health
type HealthChecker interface {
	Check(ctx context.Context) error
	Name() string
}

// Config holds the health service configuration
type Config struct {
	CheckInterval    time.Duration `yaml:"check_interval" default:"30s"`
	UnhealthyTimeout time.Duration `yaml:"unhealthy_timeout" default:"5m"`
	EnableDetailed   bool          `yaml:"enable_detailed" default:"true"`
}

// Dependencies holds all service dependencies
type Dependencies struct {
	Database HealthChecker
	Cache    HealthChecker
	Queue    HealthChecker
	External map[string]HealthChecker
}

// NewService creates a new health service with the given dependencies
func NewService(config *Config, deps Dependencies) *Service {
	if config == nil {
		config = &Config{
			CheckInterval:    30 * time.Second,
			UnhealthyTimeout: 5 * time.Minute,
			EnableDetailed:   true,
		}
	}

	service := &Service{
		status:       InternalHealthHealthy, // Use internal enum
		startTime:    time.Now(),
		lastCheck:    time.Now(),
		dependencies: make(map[string]HealthChecker),
		config:       config,
	}

	// Register core dependencies
	if deps.Database != nil {
		service.dependencies["database"] = deps.Database
	}
	if deps.Cache != nil {
		service.dependencies["cache"] = deps.Cache
	}
	if deps.Queue != nil {
		service.dependencies["queue"] = deps.Queue
	}

	// Register external dependencies
	for name, checker := range deps.External {
		service.dependencies[name] = checker
	}

	// Start background health monitoring
	go service.backgroundHealthCheck()

	return service
}

// Check implements the HealthService.Check method
func (s *Service) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	s.mu.RLock()
	currentStatus := s.status
	s.mu.RUnlock()

	// Create response using proper protobuf API (setter methods)
	response := &healthpb.HealthCheckResponse{}

	// Convert internal domain types to protobuf types
	pbStatus := convertToProtobufStatus(currentStatus)

	// This demonstrates the hybrid architecture:
	// 1. We store internal domain types (InternalHealthStatus, time.Time)
	// 2. We convert to protobuf types only at the service boundary
	// 3. We use setter methods (required for opaque protobuf API)

	// Set basic response fields using setter methods
	response.SetSummary(fmt.Sprintf("Service is %v", pbStatus))
	response.SetTotalChecks(int32(len(s.dependencies)))

	// Create metadata
	metadata := &commonpb.ResponseMetadata{}
	// metadata.SetRequestId(req.GetMetadata().GetRequestId()) // Would echo request ID if available
	response.SetMetadata(metadata)

	// TODO: Add detailed dependency status results
	// This would involve checking each dependency and converting results

	return response, nil
}

// CheckReadiness implements the HealthService.CheckReadiness method
func (s *Service) CheckReadiness(ctx context.Context, req *healthpb.ReadinessCheckRequest) (*healthpb.ReadinessCheckResponse, error) {
	// Check if all critical dependencies are healthy
	response := &healthpb.ReadinessCheckResponse{}

	for name, checker := range s.dependencies {
		if err := checker.Check(ctx); err != nil {
			// Service is not ready
			response.SetReady(false)
			response.SetReason(fmt.Sprintf("Dependency %s is not ready: %v", name, err))
			response.SetStatus(healthpb.HealthStatus_HEALTH_STATUS_UNHEALTHY)
			return response, nil
		}
	}

	// All dependencies are healthy
	response.SetReady(true)
	response.SetReason("Service is ready")
	response.SetStatus(healthpb.HealthStatus_HEALTH_STATUS_HEALTHY)
	return response, nil
}

// WatchHealth implements the HealthService.WatchHealth method for streaming health updates
// TODO: Fix the stream.Send signature issue - the grpc.ServerStreamingServer might need different parameter type
func (s *Service) WatchHealth(req *healthpb.WatchHealthRequest, stream grpc.ServerStreamingServer[*healthpb.WatchHealthResponse]) error {
	// TODO: Implement streaming health updates properly
	// Current issue: stream.Send expects **healthpb.WatchHealthResponse but we have *healthpb.WatchHealthResponse
	return nil
}

// backgroundHealthCheck runs continuous health monitoring
func (s *Service) backgroundHealthCheck() {
	ticker := time.NewTicker(s.config.CheckInterval)
	defer ticker.Stop()

	for range ticker.C {
		s.performHealthCheck()
	}
}

// performHealthCheck checks all dependencies and updates service status
func (s *Service) performHealthCheck() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	allHealthy := true

	for name, checker := range s.dependencies {
		if err := checker.Check(ctx); err != nil {
			// Log the error (in a real implementation, use structured logging)
			fmt.Printf("Health check failed for %s: %v\n", name, err)
			allHealthy = false
		}
	}

	s.mu.Lock()
	if allHealthy {
		s.status = InternalHealthHealthy
	} else {
		s.status = InternalHealthUnhealthy
	}
	s.lastCheck = time.Now()
	s.mu.Unlock()
}

// Shutdown gracefully shuts down the health service
func (s *Service) Shutdown() {
	s.mu.Lock()
	s.status = InternalHealthUnhealthy
	s.mu.Unlock()
}

// AddDependency adds a new dependency to monitor
func (s *Service) AddDependency(name string, checker HealthChecker) {
	s.mu.Lock()
	s.dependencies[name] = checker
	s.mu.Unlock()
}

// RemoveDependency removes a dependency from monitoring
func (s *Service) RemoveDependency(name string) {
	s.mu.Lock()
	delete(s.dependencies, name)
	s.mu.Unlock()
}

// HTTPHandler returns an HTTP handler for health checks (for REST API compatibility)
// This provides the same functionality as subtitle-manager's /health endpoint
func (s *Service) HTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		s.mu.RLock()
		currentStatus := s.status
		lastCheck := s.lastCheck
		dependencyCount := len(s.dependencies)
		s.mu.RUnlock()

		// Create simple health response compatible with subtitle-manager format
		response := map[string]interface{}{
			"status":    statusToString(currentStatus),
			"timestamp": lastCheck.Unix(),
			"checks": map[string]interface{}{
				"total":         dependencyCount,
				"healthy":       countHealthyDependencies(ctx, s.dependencies),
				"last_check_at": lastCheck.Format(time.RFC3339),
			},
		}

		// Set appropriate HTTP status based on health
		httpStatus := http.StatusOK
		if currentStatus == InternalHealthUnhealthy {
			httpStatus = http.StatusServiceUnavailable
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		json.NewEncoder(w).Encode(response)
	}
}

// DetailedHTTPHandler returns detailed health information for debugging
func (s *Service) DetailedHTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		s.mu.RLock()
		currentStatus := s.status
		startTime := s.startTime
		lastCheck := s.lastCheck
		s.mu.RUnlock()

		// Check each dependency individually
		dependencyResults := make(map[string]interface{})
		for name, checker := range s.dependencies {
			checkResult := map[string]interface{}{
				"name": checker.Name(),
			}

			if err := checker.Check(ctx); err != nil {
				checkResult["status"] = "unhealthy"
				checkResult["error"] = err.Error()
			} else {
				checkResult["status"] = "healthy"
			}
			dependencyResults[name] = checkResult
		}

		response := map[string]interface{}{
			"status":       statusToString(currentStatus),
			"uptime":       time.Since(startTime).String(),
			"last_check":   lastCheck.Format(time.RFC3339),
			"dependencies": dependencyResults,
			"config": map[string]interface{}{
				"check_interval": s.config.CheckInterval.String(),
				"detailed_mode":  s.config.EnableDetailed,
			},
		}

		httpStatus := http.StatusOK
		if currentStatus == InternalHealthUnhealthy {
			httpStatus = http.StatusServiceUnavailable
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			fmt.Printf("Error encoding detailed health response: %v\n", err)
		}
	}
}

// Helper functions for HTTP handlers
func statusToString(status InternalHealthStatus) string {
	switch status {
	case InternalHealthHealthy:
		return "ok"
	case InternalHealthUnhealthy:
		return "unhealthy"
	case InternalHealthDegraded:
		return "degraded"
	default:
		return "unknown"
	}
}

func countHealthyDependencies(ctx context.Context, deps map[string]HealthChecker) int {
	count := 0
	for _, checker := range deps {
		if err := checker.Check(ctx); err == nil {
			count++
		}
	}
	return count
}
