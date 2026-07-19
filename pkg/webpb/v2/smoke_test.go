// file: pkg/webpb/v2/smoke_test.go
// version: 1.0.0
// guid: 654c17f1-2906-480a-892f-6da07e851b28
// last-edited: 2026-07-19

package webv2

import "testing"

// TestPackageInitDoesNotPanic guards against a corrupted embedded protobuf
// FileDescriptor (rawDesc) in any generated file in this package: every
// .pb.go file's init() builds its descriptor unconditionally, so simply
// loading this package for testing exercises all of them. A malformed
// descriptor panics before this function body ever runs - go build/go vet
// cannot catch that class of bug, only actually running the code can.
func TestPackageInitDoesNotPanic(t *testing.T) {}
