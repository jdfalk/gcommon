// file: services/go.mod
// version: 1.0.0
// guid: e3f4a5b6-c7d8-9e0f-1a2b-3c4d5e6f7a8b

module github.com/jdfalk/gcommon/services

go 1.24

// Use local versions
replace github.com/jdfalk/gcommon => ../

replace github.com/jdfalk/gcommon/internal => ../internal/

require (
	github.com/jdfalk/gcommon/pkg/commonpb v0.0.0-20250926164202-48e80cd3f3d1
	github.com/jdfalk/gcommon/pkg/healthpb v0.0.0-20250926191523-27df7a7ed049
	github.com/spf13/viper v1.21.0
	google.golang.org/grpc v1.75.1
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
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.9-20250912141014-52f32327d4b0.1 // indirect
	github.com/spf13/cobra v1.10.1
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
)
