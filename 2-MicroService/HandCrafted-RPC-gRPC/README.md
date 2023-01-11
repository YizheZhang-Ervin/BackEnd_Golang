# Protobuf

## 编译protobuf
```
protoc --go_out=./ *.proto 会生成xx.pb.go
```

## 编译服务protobuf里的service
```
# 用grpc编译
protoc --go_out=plugins=grpc:./ *.proto
```