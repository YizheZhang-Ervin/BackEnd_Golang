# MicroService-RPC
-  GoMicro + RPC

## Dependencies
- [micro](https://github.com/micro/micro)
- [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)
- 
## 编译服务协议
```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. path/to/greeter.proto
```

## 运行
```
# 普通运行
go run server.go
go run client.go

# 微服务运行
micro run service --name helloworld
# 微服务查讯
micro call helloworld Greeter.Hello '{"name": "John"}'
```