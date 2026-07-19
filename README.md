<!-- file: README.md -->
<!-- version: 3.0.0 -->
<!-- guid: readme-gcommon-go -->

# gcommon

Protocol buffer definitions and generated Go bindings for falkcorp's shared
service types (auth, common, config, database, health, media, metrics,
organization, queue, web).

## Overview

Proto sources live on the [Buf Schema Registry](https://buf.build/falkcorp/gcommon)
(`buf.build/falkcorp/gcommon`). This repository generates and vendors the Go
bindings for that schema, and also hosts the reference gRPC service
implementations under `services/`.

The whole repository is a single Go module:

```
module github.com/falkcorp/gcommon/v2
```

There is no per-package submodule and no v1 - only the current (v2) schema
is generated. See `docs/AUDIT-2026-07-18-post-org-migration.md` for the
history of why (nested per-family modules used to exist and caused Go
module-resolution tag collisions; they were collapsed into this one module).

## Installation

```bash
go get github.com/falkcorp/gcommon/v2
```

## Usage

```go
import (
    commonv2 "github.com/falkcorp/gcommon/v2/pkg/commonpb/v2"
    healthv2 "github.com/falkcorp/gcommon/v2/pkg/healthpb/v2"
)

func main() {
    req := &healthv2.HealthCheckRequest{
        Service: "my-service",
    }
    // Use with your gRPC client:
    // client := healthv2.NewHealthServiceClient(conn)
    // resp, err := client.Check(context.Background(), req)
    _ = req
}
```

## Available Packages

| Family           | Package Path                                          |
| ---------------- | ------------------------------------------------------ |
| **auth**         | `github.com/falkcorp/gcommon/v2/pkg/authpb/v2`         |
| **common**       | `github.com/falkcorp/gcommon/v2/pkg/commonpb/v2`       |
| **config**       | `github.com/falkcorp/gcommon/v2/pkg/configpb/v2`       |
| **database**     | `github.com/falkcorp/gcommon/v2/pkg/databasepb/v2`     |
| **health**       | `github.com/falkcorp/gcommon/v2/pkg/healthpb/v2`       |
| **media**        | `github.com/falkcorp/gcommon/v2/pkg/mediapb/v2`        |
| **metrics**      | `github.com/falkcorp/gcommon/v2/pkg/metricspb/v2`      |
| **organization** | `github.com/falkcorp/gcommon/v2/pkg/organizationpb/v2` |
| **queue**        | `github.com/falkcorp/gcommon/v2/pkg/queuepb/v2`        |
| **web**          | `github.com/falkcorp/gcommon/v2/pkg/webpb/v2`          |

Service implementations (currently `health`, `auth`; the rest are tracked in
`docs/agent-tasks/`) live under `github.com/falkcorp/gcommon/v2/services/...`.

## Development

```bash
git clone --recurse-submodules https://github.com/falkcorp/gcommon.git
cd gcommon
make build
```

### Regenerating Code

Generated code is refreshed automatically via `.github/workflows/sync-protos.yml`
when the upstream BSR schema changes. To regenerate locally:

```bash
make generate
```

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
