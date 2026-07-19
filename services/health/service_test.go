// file: services/health/service_test.go
// version: 1.0.1
// guid: b2c3d4e5-f6g7-8h9i-0j1k-2l3m4n5o6p7q
// last-edited: 2026-07-19

package health

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	healthpb "github.com/falkcorp/gcommon/v2/pkg/healthpb/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// MockHealthChecker implements HealthChecker for testing
type MockHealthChecker struct {
	name    string
	healthy bool
	err     error
}

func (m *MockHealthChecker) Check(ctx context.Context) error {
	if !m.healthy {
		return m.err
	}
	return nil
}

func (m *MockHealthChecker) Name() string {
	return m.name
}

func TestNewService(t *testing.T) {
	config := &Config{
		CheckInterval:    10 * time.Second,
		UnhealthyTimeout: 1 * time.Minute,
		EnableDetailed:   true,
	}

	deps := Dependencies{
		Database: &MockHealthChecker{name: "test-db", healthy: true},
		Cache:    &MockHealthChecker{name: "test-cache", healthy: true},
		External: map[string]HealthChecker{
			"external": &MockHealthChecker{name: "external-service", healthy: true},
		},
	}

	service := NewService(config, deps)

	if service == nil {
		t.Fatal("NewService returned nil")
	}

	if service.status != InternalHealthHealthy {
		t.Errorf("Expected initial status to be healthy, got %v", service.status)
	}

	if len(service.dependencies) != 3 {
		t.Errorf("Expected 3 dependencies, got %d", len(service.dependencies))
	}
}

func TestServiceCheck(t *testing.T) {
	config := &Config{
		CheckInterval:    10 * time.Second,
		UnhealthyTimeout: 1 * time.Minute,
		EnableDetailed:   true,
	}

	deps := Dependencies{
		Database: &MockHealthChecker{name: "test-db", healthy: true},
	}

	service := NewService(config, deps)
	ctx := context.Background()

	req := &healthpb.HealthCheckRequest{}
	resp, err := service.Check(ctx, req)

	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}

	if resp == nil {
		t.Fatal("Check response is nil")
	}

	// Test response has required fields
	if resp.GetSummary() == "" {
		t.Error("Summary should not be empty")
	}

	if resp.GetTotalChecks() != 1 {
		t.Errorf("Expected 1 total check, got %d", resp.GetTotalChecks())
	}
}

func TestServiceCheckReadiness(t *testing.T) {
	healthyChecker := &MockHealthChecker{name: "healthy-service", healthy: true}
	unhealthyChecker := &MockHealthChecker{name: "unhealthy-service", healthy: false, err: context.DeadlineExceeded}

	tests := []struct {
		name      string
		deps      Dependencies
		wantReady bool
	}{
		{
			name: "all_healthy",
			deps: Dependencies{
				Database: healthyChecker,
				Cache:    healthyChecker,
			},
			wantReady: true,
		},
		{
			name: "some_unhealthy",
			deps: Dependencies{
				Database: healthyChecker,
				Cache:    unhealthyChecker,
			},
			wantReady: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &Config{
				CheckInterval:    10 * time.Second,
				UnhealthyTimeout: 1 * time.Minute,
				EnableDetailed:   true,
			}

			service := NewService(config, tt.deps)
			ctx := context.Background()

			req := &healthpb.ReadinessCheckRequest{}
			resp, err := service.CheckReadiness(ctx, req)

			if err != nil {
				t.Fatalf("CheckReadiness failed: %v", err)
			}

			if resp.GetReady() != tt.wantReady {
				t.Errorf("Expected ready=%v, got %v", tt.wantReady, resp.GetReady())
			}

			if resp.GetStatus() == healthpb.HealthStatus_HEALTH_STATUS_UNSPECIFIED {
				t.Error("Status should not be unspecified")
			}
		})
	}
}

func TestHTTPHandler(t *testing.T) {
	config := &Config{
		CheckInterval:    10 * time.Second,
		UnhealthyTimeout: 1 * time.Minute,
		EnableDetailed:   true,
	}

	deps := Dependencies{
		Database: &MockHealthChecker{name: "test-db", healthy: true},
	}

	service := NewService(config, deps)
	handler := service.HTTPHandler()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if status, ok := response["status"]; !ok || status != "ok" {
		t.Errorf("Expected status 'ok', got %v", status)
	}

	if checks, ok := response["checks"].(map[string]interface{}); !ok {
		t.Error("Response should contain 'checks' object")
	} else {
		if total, ok := checks["total"]; !ok || total != float64(1) {
			t.Errorf("Expected total checks 1, got %v", total)
		}
	}
}

func TestDetailedHTTPHandler(t *testing.T) {
	config := &Config{
		CheckInterval:    10 * time.Second,
		UnhealthyTimeout: 1 * time.Minute,
		EnableDetailed:   true,
	}

	deps := Dependencies{
		Database: &MockHealthChecker{name: "test-db", healthy: true},
		Cache:    &MockHealthChecker{name: "test-cache", healthy: false, err: context.DeadlineExceeded},
	}

	service := NewService(config, deps)
	handler := service.DetailedHTTPHandler()

	req := httptest.NewRequest(http.MethodGet, "/health/detailed", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Check for detailed response structure
	if _, ok := response["uptime"]; !ok {
		t.Error("Detailed response should contain uptime")
	}

	if deps, ok := response["dependencies"].(map[string]interface{}); !ok {
		t.Error("Detailed response should contain dependencies")
	} else {
		if len(deps) != 2 {
			t.Errorf("Expected 2 dependency results, got %d", len(deps))
		}
	}
}

func TestAddCliFlags(t *testing.T) {
	cmd := &cobra.Command{}

	// Test without prefix
	AddCliFlags(cmd, "")

	if flag := cmd.Flags().Lookup("health.check-interval"); flag == nil {
		t.Error("health.check-interval flag should be added")
	}

	if flag := cmd.Flags().Lookup("health.enable-detailed"); flag == nil {
		t.Error("health.enable-detailed flag should be added")
	}

	// Test with prefix
	cmd2 := &cobra.Command{}
	AddCliFlags(cmd2, "service")

	if flag := cmd2.Flags().Lookup("service.health.check-interval"); flag == nil {
		t.Error("service.health.check-interval flag should be added with prefix")
	}
}

func TestConfigFromViper(t *testing.T) {
	// Reset viper state
	viper.Reset()

	// Set some test values
	viper.Set("health.check_interval", "45s")
	viper.Set("health.unhealthy_timeout", "10m")
	viper.Set("health.enable_detailed", false)

	config := ConfigFromViper("")

	if config.CheckInterval != 45*time.Second {
		t.Errorf("Expected CheckInterval 45s, got %v", config.CheckInterval)
	}

	if config.UnhealthyTimeout != 10*time.Minute {
		t.Errorf("Expected UnhealthyTimeout 10m, got %v", config.UnhealthyTimeout)
	}

	if config.EnableDetailed != false {
		t.Errorf("Expected EnableDetailed false, got %v", config.EnableDetailed)
	}
}

func TestConvertToProtobufStatus(t *testing.T) {
	tests := []struct {
		internal InternalHealthStatus
		expected healthpb.HealthStatus
	}{
		{InternalHealthHealthy, healthpb.HealthStatus_HEALTH_STATUS_HEALTHY},
		{InternalHealthUnhealthy, healthpb.HealthStatus_HEALTH_STATUS_UNHEALTHY},
		{InternalHealthDegraded, healthpb.HealthStatus_HEALTH_STATUS_DEGRADED},
		{InternalHealthUnknown, healthpb.HealthStatus_HEALTH_STATUS_UNSPECIFIED},
	}

	for _, tt := range tests {
		result := convertToProtobufStatus(tt.internal)
		if result != tt.expected {
			t.Errorf("convertToProtobufStatus(%v) = %v, want %v", tt.internal, result, tt.expected)
		}
	}
}

func TestStatusToString(t *testing.T) {
	tests := []struct {
		status   InternalHealthStatus
		expected string
	}{
		{InternalHealthHealthy, "ok"},
		{InternalHealthUnhealthy, "unhealthy"},
		{InternalHealthDegraded, "degraded"},
		{InternalHealthUnknown, "unknown"},
	}

	for _, tt := range tests {
		result := statusToString(tt.status)
		if result != tt.expected {
			t.Errorf("statusToString(%v) = %v, want %v", tt.status, result, tt.expected)
		}
	}
}

func TestAddRemoveDependency(t *testing.T) {
	config := &Config{
		CheckInterval:    10 * time.Second,
		UnhealthyTimeout: 1 * time.Minute,
		EnableDetailed:   true,
	}

	service := NewService(config, Dependencies{})

	// Add dependency
	checker := &MockHealthChecker{name: "test", healthy: true}
	service.AddDependency("test", checker)

	if len(service.dependencies) != 1 {
		t.Errorf("Expected 1 dependency after adding, got %d", len(service.dependencies))
	}

	// Remove dependency
	service.RemoveDependency("test")

	if len(service.dependencies) != 0 {
		t.Errorf("Expected 0 dependencies after removing, got %d", len(service.dependencies))
	}
}

func TestShutdown(t *testing.T) {
	config := &Config{
		CheckInterval:    10 * time.Second,
		UnhealthyTimeout: 1 * time.Minute,
		EnableDetailed:   true,
	}

	service := NewService(config, Dependencies{})

	// Service should start healthy
	if service.status != InternalHealthHealthy {
		t.Errorf("Expected service to start healthy, got %v", service.status)
	}

	// Shutdown should mark as unhealthy
	service.Shutdown()

	if service.status != InternalHealthUnhealthy {
		t.Errorf("Expected service to be unhealthy after shutdown, got %v", service.status)
	}
}
