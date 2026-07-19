// file: pkg/healthpb/v2/smoke_test.go
// version: 1.0.0
// guid: 1262b12a-0620-492b-b732-8112baa5b049
// last-edited: 2026-07-19

package healthv2

import "testing"

// TestPackageInitDoesNotPanic guards against a corrupted embedded protobuf
// FileDescriptor (rawDesc) in any generated file in this package: every
// .pb.go file's init() builds its descriptor unconditionally, so simply
// loading this package for testing exercises all of them. A malformed
// descriptor panics before this function body ever runs - go build/go vet
// cannot catch that class of bug, only actually running the code can.
func TestPackageInitDoesNotPanic(t *testing.T) {}
