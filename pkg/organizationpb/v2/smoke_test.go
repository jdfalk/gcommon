// file: pkg/organizationpb/v2/smoke_test.go
// version: 1.0.0
// guid: 0235cdbe-9fab-4cfe-bd0f-56245e34f414
// last-edited: 2026-07-19

package organizationv2

import "testing"

// TestPackageInitDoesNotPanic guards against a corrupted embedded protobuf
// FileDescriptor (rawDesc) in any generated file in this package: every
// .pb.go file's init() builds its descriptor unconditionally, so simply
// loading this package for testing exercises all of them. A malformed
// descriptor panics before this function body ever runs - go build/go vet
// cannot catch that class of bug, only actually running the code can.
func TestPackageInitDoesNotPanic(t *testing.T) {}
