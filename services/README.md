<!-- file: services/README.md -->
<!-- version: 1.0.0 -->
<!-- guid: d2c3b4a5-f6e7-8a9b-0c1d-2e3f4a5b6c7d -->

# GCommon Service Implementations

This directory contains concrete implementations of all protobuf services defined in the `pkg/` directory. These implementations provide production-ready, reusable service logic that can be used across multiple applications.

## Architecture

### Service Structure

Each service follows a consistent structure:

```ascii
services/{package}/
├── service.go          # Main service implementation
├── config.go           # Service-specific configuration
├── middleware.go       # Service-specific middleware
├── handlers.go         # Business logic handlers
├── dependencies.go     # Dependency interfaces
├── example/           # Usage examples
│   └── main.go        # Standalone server example
└── README.md          # Service-specific documentation
```

### Implementation Pattern

All services follow the same implementation pattern:

1. **Interface Compliance**: Implements the generated gRPC service interface
2. **Dependency Injection**: Uses constructor injection for all dependencies
3. **Configuration**: Supports external configuration with sensible defaults
4. **Middleware**: Includes standard middleware (logging, metrics, auth, validation)
5. **Health Checking**: Supports health, readiness, and liveness probes
6. **Graceful Shutdown**: Proper resource cleanup and connection draining

### Service Dependencies

Services can depend on:

- **Database**: SQL/NoSQL database connections
- **Cache**: Redis, Memcached, or in-memory cache
- **Queue**: Message queue systems (Redis, RabbitMQ, Kafka)
- **External APIs**: Third-party service integrations
- **Other Services**: Cross-service communication

## Available Services

### Core Infrastructure

- **[health](health/)** - Health checking, readiness/liveness probes
- **[config](config/)** - Configuration management and hot-reload
- **[database](database/)** - Database operations, migrations, connection pooling
- **[queue](queue/)** - Message queuing, async processing, monitoring
- **[metrics](metrics/)** - Metrics collection, aggregation, export

### Security & Authentication

- **[auth](auth/)** - Authentication, JWT management, session handling
- **[authorization](authorization/)** - RBAC, policy enforcement, permissions

### Application Services

- **[organization](organization/)** - Multi-tenant organization management
- **[media](media/)** - Media processing, audio handling, subtitle management
- **[web](web/)** - HTTP/gRPC gateway, static assets, admin interfaces

### Operational Services

- **[cache](cache/)** - Distributed caching, invalidation strategies
- **[logging](logging/)** - Log aggregation, structured logging, analysis
- **[notification](notification/)** - Multi-channel notifications, templates
- **[workflow](workflow/)** - Process orchestration, state management

## Quick Start

### 1. Choose a Service

```go
import "github.com/falkcorp/gcommon/services/health"
```

### 2. Configure Dependencies

```go
deps := health.Dependencies{
    Database: &DatabaseHealthChecker{db: myDB},
    Cache:    &RedisHealthChecker{redis: myRedis},
    External: map[string]health.HealthChecker{
        "api_gateway": &HTTPHealthChecker{url: "http://gateway/health"},
    },
}
```

### 3. Configure Service

```go
config := &health.Config{
    CheckInterval:    30 * time.Second,
    UnhealthyTimeout: 5 * time.Minute,
    EnableDetailed:   true,
}
```

### 4. Create and Register Service

```go
healthService := health.NewService(config, deps)

server := grpc.NewServer()
healthpb.RegisterHealthServiceServer(server, healthService)
```

### 5. Start Server

```go
lis, _ := net.Listen("tcp", ":8080")
server.Serve(lis)
```

## Configuration Management

### Environment-Based Configuration

```yaml
# config.yaml
services:
  health:
    check_interval: 30s
    unhealthy_timeout: 5m
    enable_detailed: true

  database:
    host: localhost
    port: 5432
    database: myapp
    pool_size: 10
    max_idle_time: 5m
```

### Configuration Loading

```go
import "github.com/falkcorp/gcommon/internal/config"

cfg, err := config.Load("config.yaml")
if err != nil {
    log.Fatal(err)
}

healthConfig := cfg.Services.Health
dbConfig := cfg.Services.Database
```

## Middleware System

### Standard Middleware Stack

All services include a standard middleware stack:

1. **Logging**: Request/response logging with correlation IDs
2. **Metrics**: Prometheus metrics collection
3. **Authentication**: JWT token validation
4. **Authorization**: RBAC and permission checking
5. **Validation**: Request payload validation
6. **Rate Limiting**: Per-user/service rate limiting
7. **Recovery**: Panic recovery and error handling

### Custom Middleware

```go
func CustomMiddleware() grpc.UnaryServerInterceptor {
    return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
        // Custom logic here
        return handler(ctx, req)
    }
}

// Apply to service
server := grpc.NewServer(
    grpc.ChainUnaryInterceptor(
        logging.Middleware(),
        metrics.Middleware(),
        auth.Middleware(),
        CustomMiddleware(),
    ),
)
```

## Testing Strategy

### Unit Tests

Each service includes comprehensive unit tests:

```go
func TestHealthService_Check(t *testing.T) {
    deps := health.Dependencies{
        Database: &mockHealthChecker{healthy: true},
    }

    service := health.NewService(nil, deps)

    resp, err := service.Check(context.Background(), &healthpb.HealthCheckRequest{})

    assert.NoError(t, err)
    assert.Equal(t, healthpb.HealthStatus_HEALTH_STATUS_HEALTHY, resp.Status)
}
```

### Integration Tests

Full integration tests with real dependencies:

```go
func TestHealthService_Integration(t *testing.T) {
    // Start test database
    db := startTestDB(t)
    defer db.Close()

    // Create service with real dependencies
    deps := health.Dependencies{
        Database: &PostgreSQLHealthChecker{db: db},
    }

    service := health.NewService(nil, deps)

    // Test with real database connection
    resp, err := service.Check(context.Background(), &healthpb.HealthCheckRequest{})

    assert.NoError(t, err)
    assert.Equal(t, healthpb.HealthStatus_HEALTH_STATUS_HEALTHY, resp.Status)
}
```

### Load Testing

Performance testing with realistic load:

```bash
# Install ghz for gRPC load testing
go install github.com/bojand/ghz/cmd/ghz@latest

# Run load test
ghz --insecure \
  --proto health_service.proto \
  --call health.v1.HealthService.Check \
  --total 10000 \
  --concurrency 50 \
  localhost:8080
```

## Production Deployment

### Docker Container

```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o health-service ./services/health/example

FROM alpine:3.18
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/health-service /usr/local/bin/
EXPOSE 8080
CMD ["health-service"]
```

### Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: health-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: health-service
  template:
    metadata:
      labels:
        app: health-service
    spec:
      containers:
      - name: health-service
        image: gcommon/health-service:latest
        ports:
        - containerPort: 8080
          name: grpc
        livenessProbe:
          exec:
            command: ["/bin/grpc_health_probe", "-addr=:8080"]
          initialDelaySeconds: 30
        readinessProbe:
          exec:
            command: ["/bin/grpc_health_probe", "-addr=:8080"]
          initialDelaySeconds: 5
        resources:
          requests:
            memory: "64Mi"
            cpu: "50m"
          limits:
            memory: "128Mi"
            cpu: "100m"
```

### Service Mesh Integration

Services are designed to work with Istio, Linkerd, and other service mesh solutions:

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: health-service
spec:
  hosts:
  - health-service
  http:
  - match:
    - headers:
        grpc-method:
          exact: Check
    route:
    - destination:
        host: health-service
        subset: v1
    timeout: 10s
    retries:
      attempts: 3
      perTryTimeout: 3s
```

## Monitoring & Observability

### Metrics

All services export Prometheus metrics:

- **Request metrics**: Total requests, request duration, error rates
- **Business metrics**: Service-specific counters and gauges
- **Resource metrics**: Memory usage, goroutine count, database connections

### Logging

Structured logging with configurable levels:

```json
{
  "timestamp": "2024-09-26T12:00:00Z",
  "level": "info",
  "service": "health",
  "method": "Check",
  "request_id": "req-123",
  "duration": "15ms",
  "status": "success"
}
```

### Tracing

Distributed tracing with OpenTelemetry:

```go
import "go.opentelemetry.io/otel"

tracer := otel.Tracer("health-service")
ctx, span := tracer.Start(ctx, "health.Check")
defer span.End()
```

## Development Guidelines

### Code Style

- Follow Go best practices and conventions
- Use consistent error handling patterns
- Include comprehensive documentation
- Write testable code with dependency injection

### Performance

- Use connection pooling for databases
- Implement caching where appropriate
- Set reasonable timeouts and limits
- Profile performance-critical paths

### Security

- Validate all input parameters
- Use secure defaults for configurations
- Implement proper authentication/authorization
- Log security-relevant events

### Reliability

- Handle errors gracefully with proper recovery
- Implement circuit breakers for external dependencies
- Use exponential backoff for retries
- Support graceful shutdown

## Contributing

### Adding a New Service

1. Create the service directory: `services/{package}/`
2. Implement the service interface from `pkg/{package}pb/`
3. Add configuration struct with validation
4. Include comprehensive tests
5. Create usage examples
6. Update this README

### Service Template

Use the health service as a template for new services:

```bash
cp -r services/health services/myservice
# Update package names, interfaces, and logic
```

### Code Review Checklist

- [ ] Implements correct protobuf interface
- [ ] Includes comprehensive error handling
- [ ] Has configuration with defaults
- [ ] Includes unit and integration tests
- [ ] Follows consistent patterns
- [ ] Has proper documentation
- [ ] Supports graceful shutdown
- [ ] Includes middleware integration

For specific implementation details, see the individual service documentation in each service directory.
