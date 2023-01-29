# GoMicro - Project：TodoList

## 一、概述
### 1. 技术
- 微服务 gomicro/v2
- 服务发现 etcd
- 鉴权 jwt
- web框架 gin 
- 数据库 gorm + mysql
- 远程服务调用 grpc + protobuf
- 消息队列 amqp
- 服务熔断 hystrix
- 加密 crypto

### 2. 功能
- 用户注册登录(jwt鉴权)
- 新增/删除/修改/查询 备忘录

## 二、项目结构
### 1. gateway 网关部分
```
gateway/
├── pkg
│  ├── e
│  ├── logging
│  └── util
├── services
│  └── proto
├── weblib
│  ├── handlers
│  └──  middleware
└── wrappers
```
- pkg/e : 封装错误码
- pkg/logging : 日志文件
- pkg/util : 工具函数
- service/proto : 放置proto文件以及生成的pb文件
- weblib/handlers : 各个服务的接口
- weblib/middleware : http服务器的中间件
- wrappers : 放置服务熔断的配置

### 2. RabbitMQ 消息队列
```
mq-server/
├── conf
├── model
└── service
```
- conf：配置信息
- model：数据库模型
- service：服务

### 3. task & user
```
task/ & user/
├── conf
├── core
├── model
└── service
```
- conf：配置信息
- core：业务逻辑
- model：数据库模型
- service：proto文件以及各服务

## 三、部署
### 1. 运行
1. 运行RabbitMQ
2. 运行ETCD
3. 运行各模块
```
go run main.go --registry=etcd --registry_address=127.0.0.1:2379
```

### 2. 配置文件
- conf/config.ini 文件
```ini
[service]
AppMode = debug
HttpPort = :3000

[mysql]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = root
DbName = micro_todolist

[rabbitmq]
RabbitMQ = amqp
RabbitMQUser = guest
RabbitMQPassWord = guest
RabbitMQHost = localhost
RabbitMQPort = 5672