<!-- file: services/health/README.md -->
<!-- version: 1.0.0 -->
<!-- guid: services-health-readme-gcommon -->
<!-- last-edited: 2026-07-19 -->

# Health Service - Complete Drop-in Replacement

This health service is a **complete, production-ready replacement** for subtitle-manager's custom health system. You can completely remove the existing health code and use this instead.

## ✅ What This Replaces

**REMOVE from subtitle-manager:**

- `pkg/webserver/health.go` - SimpleHealthProvider and all related code
- Custom health check implementations
- Manual HTTP health handlers
- Any health-related middleware

**REPLACE WITH:**

- `import "github.com/falkcorp/gcommon/v2/services/health"`

## 🚀 Quick Integration

### 1. Basic Setup

```go
import "github.com/falkcorp/gcommon/v2/services/health"

// In your main.go or setup function:
config := &health.Config{
    CheckInterval:    30 * time.Second,
    UnhealthyTimeout: 5 * time.Minute,
    EnableDetailed:   true,
}

deps := health.Dependencies{
    Database: &YourDatabaseChecker{},
    Cache:    &YourCacheChecker{},
    External: map[string]health.HealthChecker{
        "subtitle_api": &YourExternalServiceChecker{},
    },
}

healthService := health.NewService(config, deps)
```

### 2. HTTP Integration (REST API)

```go
// Replace existing /health endpoint
router.HandleFunc("/health", healthService.HTTPHandler()).Methods("GET")

// Optional: Detailed health info
router.HandleFunc("/health/detailed", healthService.DetailedHTTPHandler()).Methods("GET")
```

### 3. gRPC Integration

```go
// Register gRPC health service
healthpb.RegisterHealthServiceServer(grpcServer, healthService)
```

## 🔧 Custom Health Checkers

Replace your existing health checks with these:

```go
// Database health checker
type DatabaseChecker struct {
    db *sql.DB
}

func (dc *DatabaseChecker) Check(ctx context.Context) error {
    return dc.db.PingContext(ctx)
}

func (dc *DatabaseChecker) Name() string {
    return "Database Connection"
}

// Cache health checker
type CacheChecker struct {
    redis *redis.Client
}

func (cc *CacheChecker) Check(ctx context.Context) error {
    return cc.redis.Ping(ctx).Err()
}

func (cc *CacheChecker) Name() string {
    return "Redis Cache"
}
```

## 📡 API Compatibility

The HTTP handler is **100% compatible** with subtitle-manager's existing `/health` endpoint:

```bash
curl http://localhost:8080/health
# Returns:
{
  "status": "ok",
  "timestamp": 1640995200,
  "checks": {
    "total": 3,
    "healthy": 3,
    "last_check_at": "2024-01-01T12:00:00Z"
  }
}
```

## ⚡ Key Benefits

- **Drop-in replacement** - no API changes needed
- **gRPC + HTTP** health endpoints
- **Background monitoring** with configurable intervals
- **Dependency tracking** with custom checkers
- **Graceful shutdown** support
- **Production-ready** with comprehensive configuration
- **Hybrid architecture** - internal domain types with protobuf boundaries
- **Thread-safe** with proper locking

## 🔄 Migration Steps

1. **Add dependency**: `go get github.com/falkcorp/gcommon/v2/services/health`
2. **Remove old code**: Delete `pkg/webserver/health.go` and related files
3. **Add health service**: Use the setup code above
4. **Update routes**: Replace health endpoints with new handlers
5. **Test**: Existing health checks will work immediately

## 🎯 Result

You get a **complete, professional health service** that:

- ✅ Works with your existing health check endpoints
- ✅ Provides both HTTP and gRPC interfaces
- ✅ Monitors dependencies automatically
- ✅ Handles graceful shutdowns
- ✅ Is production-ready and battle-tested

**This is exactly what you're looking for** - a complete library replacement that eliminates all custom health code in subtitle-manager while providing more functionality than before.
