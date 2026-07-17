module github.com/falkcorp/gcommon

go 1.25.0

// Replace directives for local sub-modules
replace github.com/falkcorp/gcommon/internal => ./internal

replace github.com/falkcorp/gcommon/services => ./services

replace github.com/falkcorp/gcommon/services/health => ./services/health

replace github.com/falkcorp/gcommon/services/auth => ./services/auth

replace github.com/falkcorp/gcommon/pkg/authpb/v2 => ./pkg/authpb/v2

require (
	github.com/falkcorp/gcommon/internal v0.0.0-20251003134307-5cabf522c911
	google.golang.org/grpc v1.79.3
)

require (
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
