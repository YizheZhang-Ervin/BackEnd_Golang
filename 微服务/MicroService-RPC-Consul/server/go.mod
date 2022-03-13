module dry_server_plus

go 1.14

require (
	dry_base_plus v0.0.0
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.7.0
	github.com/micro/go-plugins/registry/consul/v2 v2.5.0
)

replace dry_base_plus => ../base
