# Demo
- GoMicro

## Web as MicroService
```
# 启动服务
micro service --name helloworld --endpoint http://localhost:9090 go run helloworld.go
# 查询服务
micro call -o raw helloworld /
```
## Text as MicroService
```
# 运行服务
micro service --name helloworld --endpoint file:///tmp/helloworld.txt
# 获取文件
micro call -o raw helloworld .
```

## Shell as MicroService
```
# 远程执行脚本或命令
micro service --name helloworld --endpoint exec:///tmp/hellworld.sh
micro call -o raw helloworld .
```