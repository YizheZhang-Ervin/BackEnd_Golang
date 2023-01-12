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
consul agent -dev
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