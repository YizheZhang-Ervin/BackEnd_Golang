# MyMiddle

## 命令
```
# golang
go mod init xxMod
go mod tidy
go build
go run

# etcd
go get go.etcd.io/etcd/client/v3
etcdctl--endpoints=http://127.0.0.1:2379 put xxKey"xxVal"
etcdctl--endpoints=http://127.0.0.1:2379 get xxKey
```