// file: pkg/commonpb/v2/smoke_test.go
// version: 1.0.0
// guid: b5d9df88-e5ad-465e-8f62-a890a59c58cf
// last-edited: 2026-07-19

package commonv2

import "testing"

// TestPackageInitDoesNotPanic guards against a corrupted embedded protobuf
// FileDescriptor (rawDesc) in any generated file in this package: every
// .pb.go file's init() builds its descriptor unconditionally, so simply
// loading this package for testing exercises all of them. A malformed
// descriptor panics before this function body ever runs - go build/go vet
// cannot catch that class of bug, only actually running the code can.
func TestPackageInitDoesNotPanic(t *testing.T) {}
