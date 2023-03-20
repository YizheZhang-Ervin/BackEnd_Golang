# Go Micro v3

## 安装
```
go install github.com/micro/micro/v3@latest
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v3/cmd/protoc-gen-micro
```

## 启动micro
```
micro server
micro login  # username ‘admin’ and password ‘micro’
micro services
```

## 调用
```
micro helloworld call --name=Jane
curl "http://localhost:8080/helloworld?name=John"
```

## 创服务
```
micro new helloworld
cd helloworld
make init
go mod vendor
make proto
```

## 启服务
```
micro run .
micro update .
```

## 监控
```
micro status
micro logs helloworld
```