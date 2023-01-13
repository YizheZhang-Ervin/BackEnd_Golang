# Protobuf
```
# 编译protobuf
protoc --go_out=./ *.proto 会生成xx.pb.go

# 编译服务protobuf里的service (用grpc编译)
protoc --go_out=plugins=grpc:./ *.proto
```

# grpc
```
# 安装
go get -u -v google.golang.org/grpc
或
git clone 依赖再go install
或
unzip离线包再go install
```

# module,import,folder,package
```
module模块
import 文件绝对路径
folder一般和package同名
package main入口
```

# 微服务
```
go-micro核心库
micro运行环境&命令&创建微服务空项目
go-plugins插件
protoc-gen-micrp生成micro相关代码
```

# consul
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

# TODO修改代码
```
gomicro-client 源于test77
gomicro-server 源于test66
gomicro-gin 源于test66web
gomicro-gin-gorm 源于web
microservice1 源于bj38web
microservice2 源于ihomebj5q
```