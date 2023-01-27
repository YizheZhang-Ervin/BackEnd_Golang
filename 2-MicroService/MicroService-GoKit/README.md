# GoKit

# 1. 三层结构
Transport：http/grpc/thrift
Endpoint：request/response
Service：业务类

# 2. 代码
```
服务端 TES：路由、http方法、注册反注册consul、统一异常处理、日志、jwt、限流
客户端 TES-client：客户端调用服务端、注册反注册consul、负载均衡调用(随机/轮询)、熔断
限流 Rate
容错、降级、熔断 Hystrix
认证 JWT
```