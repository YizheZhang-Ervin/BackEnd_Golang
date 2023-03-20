module dry_client_plus

go 1.14

require (
	dry_base_plus v0.0.0
	github.com/micro/go-micro/v2 v2.7.0
	github.com/micro/go-plugins/registry/consul/v2 v2.5.0
)

replace dry_base_plus => ../base
