# GRPC - Project：TodoList

## 一、概述
### 1. 技术
- 服务发现 etcd
- web框架 gin 
- 远程服务调用 grpc + protobuf
- 数据库 gorm + mysql
- 鉴权 jwt
- 日志 logrus
- 配置库 viper

## 二、项目结构
### 1. gateway 网关部分
```
api-gateway/
├── cmd                   // 启动入口
├── config                // 配置文件
├── discovery             // etcd服务注册、keep-alive、获取服务信息等等
├── internal              // 业务逻辑（不对外暴露）
│   ├── handler           // 视图层
│   └── service           // 服务层
│       └──pb             // 放置生成的pb文件
├── logs                  // 放置打印日志模块
├── middleware            // 中间件
├── pkg                   // 各种包
│   ├── e                 // 统一错误状态码
│   ├── res               // 统一response接口返回
│   └── util              // 各种工具、JWT、Logger等等..
├── routes                // http路由模块
└── wrappers              // 各个服务之间的熔断降级
```

### 2. user && task 用户与任务模块
```
user/
├── cmd                   // 启动入口
├── config                // 配置文件
├── discovery             // etcd服务注册、keep-alive、获取服务信息等等
├── internal              // 业务逻辑（不对外暴露）
│   ├── handler           // 视图层
│   ├── cache             // 缓存模块
│   ├── repository        // 持久层
│   └── service           // 服务层
│       └──pb             // 放置生成的pb文件
├── logs                  // 放置打印日志模块
└── pkg                   // 各种包
    ├── e                 // 统一错误状态码
    ├── res               // 统一response接口返回
    └── util              // 各种工具、JWT、Logger等等..
```

## 三、部署
### 1. 运行
1. 运行ETCD
2. 运行各模块
```
go mod tidy
go run main.go
```

### 2. 配置文件
- config/config.yml
```yaml
server:
# 模块
  domain: user
  # 模块名称
  version: 1.0
  # 模块版本
  grpcAddress: "127.0.0.1:10001"
  # grpc地址

datasource:
# mysql数据源
  driverName: mysqlMaster
  host: 127.0.0.1
  port: 3306
  database: basicInfo
  # 数据库名
  username: root
  password: root
  charset: utf8mb4

etcd:
# etcd 配置
  address: 127.0.0.1:2379

redis:
# redis 配置
  address: 127.0.0.1:6379
  password:
```