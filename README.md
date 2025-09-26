<!-- file: README.md -->
<!-- version: 2.0.0 -->
<!-- guid: readme-gcommon-go -->

# gcommon-go

Go SDK for gcommon protocol buffers with generated Go bindings

## Overview

This repository provides Go bindings for the [gcommon](https://github.com/jdfalk/gcommon) protocol buffer definitions via the [Buf Schema Registry](https://buf.build/jdfalk/gcommon). It includes all generated Go code and provides a stable Go module for consuming gcommon services.

## Features

- **BSR Integration**: Uses buf.build/jdfalk/gcommon for protocol buffer definitions
- **Generated Go Code**: Pre-generated Go bindings for all gcommon services
- **Versioned Releases**: Follows semantic versioning for compatibility
- **Module Support**: Proper Go module with github.com/jdfalk/gcommon path
- **Documentation**: Comprehensive API documentation

## Installation

```bash
go get github.com/jdfalk/gcommon
```

## Usage

### Import the Generated Code

```go
import (
    "github.com/jdfalk/gcommon/common/v1"
    "github.com/jdfalk/gcommon/health/v1"
    "github.com/jdfalk/gcommon/metrics/v1"
    // ... other modules
)
```

### Example Usage

```go
package main

import (
    "context"
    "log"

    commonv1 "github.com/jdfalk/gcommon/common/v1"
    healthv1 "github.com/jdfalk/gcommon/health/v1"
)

func main() {
    // Create a health check request
    req := &healthv1.HealthCheckRequest{
        Service: "my-service",
    }

    // Use with your gRPC client
    // client := healthv1.NewHealthServiceClient(conn)
    // resp, err := client.Check(context.Background(), req)
}

```

## Available Modules

The Go SDK includes bindings for all 9 gcommon modules:

| Module           | Package Path                                | Description                      |
| ---------------- | ------------------------------------------- | -------------------------------- |
| **common**       | `github.com/jdfalk/gcommon/common/v1`       | Shared types, errors, pagination |
| **config**       | `github.com/jdfalk/gcommon/config/v1`       | Configuration management         |
| **database**     | `github.com/jdfalk/gcommon/database/v1`     | Database operations              |
| **health**       | `github.com/jdfalk/gcommon/health/v1`       | Health checks and monitoring     |
| **media**        | `github.com/jdfalk/gcommon/media/v1`        | Media processing                 |
| **metrics**      | `github.com/jdfalk/gcommon/metrics/v1`      | System metrics                   |
| **organization** | `github.com/jdfalk/gcommon/organization/v1` | Organization management          |
| **queue**        | `github.com/jdfalk/gcommon/queue/v1`        | Message queuing                  |
| **web**          | `github.com/jdfalk/gcommon/web/v1`          | Web services                     |

## Development

### Building from Source

```bash
git clone https://github.com/jdfalk/gcommon-go.git
cd gcommon-go
make build
```

### Regenerating Code

The generated code is updated automatically via GitHub Actions when the upstream gcommon repository changes. To regenerate locally:

```bash
buf generate
```

## Related Repositories

- **[gcommon](https://github.com/jdfalk/gcommon)** - Protocol buffer definitions (BSR: buf.build/jdfalk/gcommon)
- **[gcommon-py](https://github.com/jdfalk/gcommon-py)** - Python SDK

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
