// file: services/auth/go.mod
// version: 1.0.0
// guid: h3i4j5k6-l7m8-n9o0-p1q2-r3s4t5u6v7w8

module github.com/jdfalk/gcommon/services/auth

go 1.24.0

toolchain go1.24.7

require (
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/jdfalk/gcommon/pkg/authpb v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.16.0
	google.golang.org/grpc v1.65.0
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/jdfalk/gcommon/pkg/authpb => ../../pkg/authpb
