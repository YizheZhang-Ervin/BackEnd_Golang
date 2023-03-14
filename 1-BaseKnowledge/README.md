# Golang

## 0.GO目录
```
Go 语言环境安装
Go 语言结构
Go 语言基础语法
Go 语言数据类型
Go 语言变量
Go 语言常量
Go 语言运算符
Go 语言条件语句
Go 语言循环语句
Go 语言函数
Go 语言变量作用域
Go 语言数组
Go 语言指针
Go 语言结构体
Go 语言切片(Slice)
Go 语言范围(Range)
Go 语言Map(集合)
Go 语言递归函数
Go 语言类型转换
Go 语言接口
Go 错误处理
Go 并发
Go 语言开发工具(vscode/goland/liteIDE/Eclipse)
```

## 1.环境准备
```
godep
go mod
```

## 2.基础
```
变量常量(标识符、关键字、变量、常量)
数据类型(基本数据类型、类型转换)
运算符(算术、关系、逻辑、位、赋值)
流程控制(ifelse、for、forrange、switchcase、goto、break、continue)
数组(定义、初始化、遍历、多维、值类型)
切片(定义、赋值拷贝、遍历、加元素、扩容、复制、删除)
map(定义、使用、判键存在、遍历、删除、顺序遍历、切片map)
函数(参数、调用、返回值、作用域、函数类型变量、高阶函数、匿名、闭包、defer、内置函数)
指针(地址、类型、取值、new和make)
结构体(类型别名及自定义类型、结构体[定义、实例化、初始化、内存布局、构造函数、方法和接收者、任意类型添加方法、结构体的匿名字段、嵌套结构体、嵌套匿名字段、嵌套结构体的字段名冲突、结构体的“继承”、结构体字段的可见性、结构体与JSON序列化、结构体标签（Tag）、结构体和方法补充知识点])
包(定义、引入、初始化、标识符可见性)
接口(类型、值接收者及指针接收者、类型及接口、接口组合、空接口、接口值、类型断言)
error接口(创建错误、错误结构体类型)
反射(reflect[typeof、valueof]、结构体反射)
并发(goroutine、channel、select多路复用、并发安全及锁、原子操作)
网络编程(TCP、黏包、UDP)
单元测试 (基础[测试函数、基准测试、setup及teardown、示例函数]、网络测试、mysql&redis、mock接口、monkey打桩、goconvey、可测试的代码)
```

## 3.库
```
fmt
time
flag
log
文件os
strconv
net/http
context
gob/msgpack
http/template
三方logrus日志库
三方zap
三方gopsutil系统性能库
```

## 4.中间件
```
mysql (sqlx+GORM)
redis
mongoDB
influxDB
etcd
kafka
elasticSearch
NSQ
RabbitMQ
```

## 5.web
```
template
gin (JWT、源码、zap日志、路由、validator)
air热加载
优雅关机重启
docker
cookie&session
swagger
压测
限流
部署
```

## 6.其他
```
JSON
函数式选项模式
单例模式
结构体转map[string]interface{}的若干方法
Go语言配置管理神器Viper
protobuf
gRPC
pprof性能调优
makefile
在select语句中实现优先级
Go语言结构体的内存布局
```