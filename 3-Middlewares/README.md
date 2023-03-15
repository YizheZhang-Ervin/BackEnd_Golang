# MyMiddle

## 中间件
```
框架gin
数据库mysql，tidb(暂无)
键值对数据库redis，etcd
文档数据库mongoDB
对象存储minio，ceph(暂无)
消息队列kafka
弹性搜索elasticSearch
时序数据库influxDB
调度服务zookeeper(暂无)
```

## 调度
```
数据库MongoDB operator(暂无)
缓存redis operator(暂无)
大数据spark operator(暂无)
消息队列kafka operator(暂无)
调度服务zookeeper operator(暂无)
```

## 命令
```
# golang
go mod init myMiddle
go mod tidy
go build
go run

# gin
go get -u github.com/gin-gonic/gin

# etcd
go get go.etcd.io/etcd/client/v3
etcdctl--endpoints=http://127.0.0.1:2379 put xxKey"xxVal"
etcdctl--endpoints=http://127.0.0.1:2379 get xxKey

# redis
go get github.com/redis/go-redis/v9

# mysql
go get -u github.com/go-sql-driver/mysql

# minio
go get github.com/minio/minio-go/v7

# kafka
go get github.com/segmentio/kafka-go

# mongoDB
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
```