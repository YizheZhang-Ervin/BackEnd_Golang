# Linux
```
tar -C /usr/local -xzf go1.4.linux-amd64.tar.gz

cat >> ~/.bash_profile <<EOF
export PATH=$PATH:/usr/local/go/bin
EOF

source ~/.bash_profile
```

# Windows
```
# 把go/bin配在PATH中

# 其他环境变量
go env -w GOPATH=D:\GitRepository\TMPL_BackupPrograms\#Golang
go env -w GOMODCACHE=D:\GitRepository\TMPL_BackupPrograms\#Golang\pkg\mod
go env -w GO111MODULE=on
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
或
go env -w GOPROXY=https://goproxy.cn
```

# GO目录
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