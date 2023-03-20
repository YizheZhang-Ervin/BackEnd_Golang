module helloworld

go 1.18

require (
	github.com/golang/protobuf v1.5.3
	github.com/micro/micro/v3 v3.16.0
	google.golang.org/protobuf v1.30.0
)

require google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
