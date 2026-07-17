// file: pkg/databasepb/go.mod
// version: 1.0.1
// guid: go-mod-databasepb-v1

// Deprecated: This module is deprecated. Use github.com/falkcorp/gcommon/pkg/databasepb/v2 instead.
// The v2 module provides enhanced functionality, additional methods, and improved protobuf definitions.
module github.com/falkcorp/gcommon/pkg/databasepb

go 1.25.0

replace github.com/falkcorp/gcommon/pkg/commonpb => ../commonpb

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.10-20250912141014-52f32327d4b0.1
	github.com/falkcorp/gcommon/pkg/commonpb v0.0.0-20251003134307-5cabf522c911
	google.golang.org/grpc v1.79.3
	google.golang.org/protobuf v1.36.10
)

require (
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
)
