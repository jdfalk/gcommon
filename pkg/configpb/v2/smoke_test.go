// file: pkg/configpb/v2/smoke_test.go
// version: 1.0.0
// guid: 094ad9b5-f063-4aab-8c48-fef45d800acf
// last-edited: 2026-07-19

package configv2

import "testing"

// TestPackageInitDoesNotPanic guards against a corrupted embedded protobuf
// FileDescriptor (rawDesc) in any generated file in this package: every
// .pb.go file's init() builds its descriptor unconditionally, so simply
// loading this package for testing exercises all of them. A malformed
// descriptor panics before this function body ever runs - go build/go vet
// cannot catch that class of bug, only actually running the code can.
func TestPackageInitDoesNotPanic(t *testing.T) {}
