module github.com/jdfalk/gcommon

go 1.24

// Replace directives for local sub-modules
replace github.com/jdfalk/gcommon/internal => ./internal

replace github.com/jdfalk/gcommon/services => ./services

replace github.com/jdfalk/gcommon/services/health => ./services/health

replace github.com/jdfalk/gcommon/services/auth => ./services/auth

replace github.com/jdfalk/gcommon/pkg/authpb/v2 => ./pkg/authpb/v2

require (
	github.com/jdfalk/gcommon/internal v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.75.1
)

require (
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
