# MicroService - gozero

- go zero & etcd
- mysql
- redis
- jwt
- 中间件/error/模板
- log
- prometheus
- jaeger
- transaction
- 命令

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
#### (1) 生成框架
```
mkdir micro
cd micro
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
删除user/的内容(etc/internal等)，把rpc文件夹内生成的内容移出到user/，调整internal内依赖
cd ..
go mod tidy
在user/internal/logic/getuserlogic.go加入处理逻辑
```

#### (3) order
```
在order/order.api中增加接口
cd ../order
goctl api go -api order.api -dir ./gen
把gen/里生成的内容移到order/里面，删除gen/
在order/internal/config/config.go增加userrpc配置
在order/etc/order-api.yaml修改配置
在order/internal/svc/增加svc对user的rpc调用
在order/internal/logic/orderlogic.go加入处理逻辑
```

#### (4) 启动
```
增加docker-compose.yml
增加.env
docker-compose up -d
cd user
go run user.go
cd order
go run order.go
访问http://localhost:8888/api/order/get/1
```

## 2. Mysql
```
cd user/internal/model
goctl model mysql ddl -src user.sql -dir . -c
拆分生成的文件到model和repo
新增dao和database
修改rpc/user.proto，生成代码，新内容合并入user/，编写user/internal/logic/getuserlogic逻辑
cd micro-middleware
goctl api new userapi
cd userapi
go mod init userapi
cd ..
go work use userapi
cd userapi
go mod tidy
提取user内容到rpc-common
go work use rpc-common
go run userapi.go
访问http://localhost:8888/register,请求体{name,gender}
```

## 3. Redis
```
访问http://localhost:8888/user/get/6
```

## 4. Jwt
```
修改api文件
goctl api go -api userapi.api -dir ./gen后合并
访问http://localhost:8888/login
访问http://localhost:8888/api/order/get/1
```

## 5. 中间件
```
修改api文件
goctl api go -api userapi.api -dir ./gen后合并
访问http://localhost:8888/login
访问http://localhost:8888/api/order/get/1
```

## 6. Error
```
修改userapi.go
访问http://localhost:8888/api/order/get/1
```

## 7. Template
```
goctl template init修改模板
```

## 8. Log
```
增加zapx
访问http://localhost:8888/api/order/get/1
```

## 9. prometheus
```
访问http://localhost:9090/targets?search=
```

## 10. jaeger
```
访问http://localhost:16686/
```

## 11. transaction
```
dtm with go-zero
```

## 12. Other
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