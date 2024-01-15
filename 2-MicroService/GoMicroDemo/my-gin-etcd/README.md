# Gin Etcd

## 一、项目
- my-gin-etcd
    - gin + etcd
    - viper

## 二、技术栈
- 配置
    - viper
    - gopkg.in/ini.v1
- 服务治理
    - service mesh(istio&envoy)【暂无】
    - serverless(kubeless)【暂无】
- 微服务 
    - gomicro
    - gokit【暂无】
    - gozero
- 服务发现
    - etcd/k8s
- 服务熔断/负载均衡
    - hystrix
    - rate【暂无】
- 网关
    - api gateway
    - web框架 gin
- 中间件
    - mysql(gorm&validate)
    - redis(redigo)
    - rabbitMQ
    - elasticSearch【暂无】
    - fastdfs / ceph【暂无】
    - tidb【暂无】
    - spark【暂无】
    - zookeeper【暂无】
    - 各中间件operator【暂无】
- 会话状态 cookie&session
- 身份认证 jwt
- 跨域 cors
- 全局异常
- 日志

## 三、命令
```
# 1. golang
go clean --modcache
go mod init xxMod
go mod tidy
go mod vendor
go get xx/xx
go build
go run

# 2. etcd 
etcdctl--endpoints=http://127.0.0.1:2379 put xxKey"xxVal"
etcdctl--endpoints=http://127.0.0.1:2379 get xxKey
```
