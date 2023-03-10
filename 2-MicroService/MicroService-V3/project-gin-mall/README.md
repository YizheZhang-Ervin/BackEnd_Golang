# gin-mall

## 一、概述
### 1. 技术
- web框架 gin
- 数据库 mysql + gorm + dbresolver读写分离
- 缓存 redis (存储商品的浏览次数)
- 鉴权 jwt
- 加密 crypto
- 日志 logrus
- ELK

### 2. 功能
- 用户注册登录(JWT-Go鉴权)
- 用户基本信息修改，解绑定邮箱，修改密码
- 商品的发布，浏览等
- 购物车的加入，删除，浏览等
- 订单的创建，删除，支付等
- 地址的增加，删除，修改等
- 各个商品的浏览次数，以及部分种类商品的排行
- 设置了支付密码，对用户的金额进行了对称加密AES(数据脱敏)
- 支持事务，支付过程发送错误进行回退处理
- 可以将图片上传到对象存储，也可以切换分支上传到本地static目录下
- 添加ELK体系，方便日志查看和管理

## 二、项目结构
```
gin-mall/
├── api
├── cache
├── conf
├── dao
├── doc
├── middleware
├── model
├── pkg
│  ├── e
│  └── util
├── routes
├── serializer
└── service
```
- api : 用于定义接口函数
- cache : 放置redis缓存
- conf : 用于存储配置文件
- dao : 对持久层进行操作
- doc : 存放接口文档
- middleware : 应用中间件
- model : 应用数据库模型
- pkg/e : 封装错误码
- pkg/util : 工具函数
- routes : 路由逻辑处理
- serializer : 将数据序列化为 json 的函数
- service : 接口函数的实现

## 三、部署
### 1. 运行
1. 运行ES、kibana (本地 or docker-compose up)
2. 运行各模块
```
go mod tidy
go run main.go
```

### 2. 配置文件
- conf/config.ini文件配置

```ini
#debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000

[mysql]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = root
DbName = 

[redis]
RedisDb = redis
RedisAddr = 127.0.0.1:6379
RedisPw =
RedisDbName = 

[qiniu]
AccessKey = 
SerectKey = 
Bucket = 
QiniuServer = 

[email]
ValidEmail=http://localhost:8080/#/vaild/email/
SmtpHost=smtp.qq.com
SmtpEmail=
SmtpPass=
#SMTP服务的通行证

[es]
EsHost = 127.0.0.1
EsPort = 9200
EsIndex = mylog
```