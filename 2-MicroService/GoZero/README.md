# MicroService - gozero

## 1. Base

### 1.1 安装goctl
```
go install github.com/zeromicro/go-zero/tools/goctl@v1.6.1
```

### 1.2 安装rpc协议生成工具
```
goctl env check -i -f --verbose
```

### 1.3 helloworld
```
cd 1-base
goctl api new hello
cd hello
go mod tidy
在internal/logic/hellologic.go加入处理逻辑
go run hello.go访问http://localhost:8888/from/me
```

### 1.4 microhelloworld
#### (1) basic
```
cd 1-base
mkdir microhelloworld
cd microhelloworld
goctl api new order
goctl api new user
go work init
go work use order
go work use user
```

#### (2) user
```
cd user/rpc
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
cd ..
go mod tidy
在user/rpc/internal/logic/userlogic.go加入处理逻辑
```

#### (3) order
```
在order/order.api中增加接口
cd ../order
goctl api go -api order.api -dir ./gen
在order/gen/internal/config/config.go增加userrpc配置
在order/gen/etc/order-api.yaml修改配置
在order/gen/internal/svc/增加svc对user的rpc调用
在order/gen/internal/logic/orderlogic.go加入处理逻辑
```

#### (4) 启动
```
增加docker-compose.yml
增加.env
docker-compose up -d
cd user/rpc
go run user.go
cd order/gen
go run order.go
访问// http://localhost:8888/api/order/get/1
```

## 2. Other
```
# 修改模板
goctl template init

# 生成文件格式
goctl api go -api userapi.api -dir ./gen
goctl api go -api userapi.api -dir ./gen -style go_zero

# 生成proto文件
goctl rpc template -o=user.proto

# 生成rpc服务代码
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

# model
goctl model mysql ddl -src="./*.sql" -dir="./sql/model" -c

# dockerfile
goctl docker -go hello.go

# k8s
goctl kube deploy -name redis -namespace adhoc -image redis:6-alpine -o redis.yaml -port 6379

# 日志zap
writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)

# 监控prometheus
访问：http://localhost:9090/targets?search=

# 链路追踪jaeger
访问：http://localhost:16686/

# 分布式事务dtm
git clone https://github.com/dtm-labs/dtm.git
go run main.go -c conf.yml
go get github.com/dtm-labs/dtm
go get github.com/dtm-labs/driver-gozero
```