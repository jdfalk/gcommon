// file: pkg/queuepb/v2/go.mod
// version: 1.0.0
// guid: go-mod-queuepb-v2

module github.com/falkcorp/gcommon/pkg/queuepb/v2

go 1.25.0

replace github.com/falkcorp/gcommon/pkg/commonpb/v2 => ../../commonpb/v2

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.10-20250912141014-52f32327d4b0.1
	github.com/falkcorp/gcommon/pkg/commonpb/v2 v2.0.0-20251003134307-5cabf522c911
	google.golang.org/grpc v1.79.3
	google.golang.org/protobuf v1.36.10
)

require (
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
)
