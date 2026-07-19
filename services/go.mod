// file: services/go.mod
// version: 1.0.0
// guid: e3f4a5b6-c7d8-9e0f-1a2b-3c4d5e6f7a8b

module github.com/falkcorp/gcommon/services

go 1.25.0

// Use local versions
replace github.com/falkcorp/gcommon => ../

replace github.com/falkcorp/gcommon/internal => ../internal/

replace github.com/falkcorp/gcommon/pkg/commonpb/v2 => ../pkg/commonpb/v2

replace github.com/falkcorp/gcommon/pkg/healthpb/v2 => ../pkg/healthpb/v2

require (
	github.com/falkcorp/gcommon/pkg/commonpb/v2 v2.0.0-20251003134307-5cabf522c911
	github.com/falkcorp/gcommon/pkg/healthpb/v2 v2.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.21.0
	google.golang.org/grpc v1.79.3
)

require (
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/sagikazarmark/locafero v0.11.0 // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.10-20250912141014-52f32327d4b0.1 // indirect
	github.com/spf13/cobra v1.10.1
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
