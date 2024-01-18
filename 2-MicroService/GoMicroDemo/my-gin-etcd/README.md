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
# golang
go mod init xxx
go mod tidy
go mod edit -replace github.com/coreos/bbolt@v1.3.8=go.etcd.io/bbolt@v1.3.8
go mod edit -replace google.golang.org/grpc=google.golang.org/grpc@v1.26.0
go build
go run xx.go

# etcd
etcdctl--endpoints=http://127.0.0.1:2379 put xxKey"xxVal"
etcdctl--endpoints=http://127.0.0.1:2379 get xxKey
```

## 四、测试
```
Zm9vOmJhcg是base64("foo:bar")
curl -X POST http://localhost:8080/admin \
-H 'authorization: Basic Zm9vOmJhcg==' \
-H 'content-type: application/json' \
-d '{"value":"bar"}'
```