# Go 微服务
```
# RPC
- my-rpc
- my-grpc
- my-grpc-consul

# GoMicro
- my-gomicro-server
- my-gomicro-client
- my-gomicro-gin
    - Gin

# Example
- my-microservice1
    - web 客户端
        - gorm
    - service 微服务
        - getcaptcha 图片验证码
        - user 用户+短信验证码
- my-microservice2
```

# 1. Protobuf
```
# 编译protobuf
protoc --go_out=./ *.proto 会生成xx.pb.go

# 编译服务protobuf里的service (用grpc编译)
protoc --go_out=plugins=grpc:./ *.proto
```

# 2. grpc
```
# 安装
go get -u -v google.golang.org/grpc
或
git clone 依赖再go install
或
unzip离线包再go install
```

# 3. module,import,folder,package
```
module模块
import 文件绝对路径
folder一般和package同名
package main入口
```

# 4. 微服务
```
go-micro核心库
micro运行环境&命令&创建微服务空项目
go-plugins插件
protoc-gen-micrp生成micro相关代码
```

# 5. consul
```
consul -h
consul agent 
    -bind地址,-http-port端口,-client客户端,-config-dir存服务信息,
    -data-dir存机器信息,-dev开发模式,-node服务发现名,-rejoin加入,-server服务端,-ui页面
consul members 集群成员
consul info 信息
consul leave 优雅关闭
注册服务/etc/consul.d/XX.json里面{service:{name,tags,port}}
查询服务ip:port/v1/catalog/service/xx
健康检查/etc/cnsul.d/xx.json里面{service:{name,tags,address,port,check:{id,name,http,interval,timeout}}}
```

# 6. Go Micro
## 通用
```
# protobuf
https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip

# make(mingw)
https://osdn.dl.osdn.net/mingw/68260/mingw-get-setup.exe

# protoc的micro插件
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v3/cmd/protoc-gen-micro
```

## V1
```
# 安装
go get -u -v github.com/micro/go-micro
go get -u -v github.com/micro/micro
或
docker pull microhq/micro

# 命令
micro new --namespace包名 --type微服务类型(api/fnc/srv/web) --fqdn服务正式完整定义 --alias别名
如micro new --type srv xxName
使用Makefile：make proto
```
## V2
```
# 安装(go要1.15以下)
go install github.com/micro/micro/v2

# 命令
micro new --type service xxName
micro new --type web xxName
```

## V3
```
# 安装
go install github.com/micro/micro/v3@latest 或 docker pull ghcr.io/micro/micro:latest

# 命令
micro new xx
micro web xx
make proto
micro server
sudo docker run -p 8080:8080 -p 8081:8081 ghcr.io/micro/micro:latest server
micro login # admin-micro
```

## V4
```
# 安装
go install github.com/go-micro/cli/cmd/go-micro@latest

# 命令
go-micro new service xx
go-micro new client xx
```

# 7. http概念
```
# 概念
路由器资源分发
路由请求分析：service.HandleFunc("/xx",handler.xx)
go的web框架:gin,beego,echo,iris

# Gin框架
路径参数: ctx.Param(xx)
问号查询参数: ctx.Query(xx)或ctx.DefaultQuery(xx,defaultVal)
表单数据: ctx.PostForm(xx)或ctx.DefaultPostForm(xx,defaultVal)
请求体: 
    ctx.Bind()或ctx.BindJSON(&body)
    或ctx.GetRawData()及json.Unmarshal(bytes,&jsonMap)
    或ioutil.ReadAll(ctx.Request.Body)及json.Unmarshal(bytes,&jsonMap)
    或json.NewDecoder(ctx.Request.Body).Decode(&jsonMap)
```

# 8. redis
```
# 服务启动
/etc/redis/redis.conf改bind地址
port：6379
redis-server /etc/redis/redis.conf

# 客户端使用
redis-cli -h xx -p xx
keys *
flushall
set key value
get key

# go编程
redigo：连库、操作库、回复助手(把返回值转为特定类型)
```

# 9. Cookie & Session
```
cookie客户端
session服务端（k:sessionId，v:sessionValue）
```

# 10. 中间件
```
对之后的路由都生效
承上启下用于两个模块之间的功能软件(路由-中间件-控制器)
gin中间件：gin.HandlerFunc
- ctx.next跳过当前中间件剩余内容执行下一个中间件。前面的顺序调用，后面的逆序调用
- abort只执行当前中间件，阻止执行下一个中间件
- return终止执行当前中间件剩余内容，执行下一个中间件
```

# 11. 综合
```
go micro、gorm+validate、redigo、gin、熔断器hystrix
网关http api、网关grpc、consul、etcd、限流rate、jwt、go-kit
cookie&session、nginx、micro web、micro registry
```