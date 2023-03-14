# Golang 安装&配置

1. 安装 - Linux
```
tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz

cat >> ~/.bash_profile <<EOF
export PATH=$PATH:/usr/local/go/bin
EOF

source ~/.bash_profile
```

2. 安装 - Windows
```
# 把go/bin配在PATH中

# 其他环境变量(goproxy配了之后再装vscode插件)
go env -w GOPATH=D:\GoPathRepository (要和go mod路径不一样)
go env -w GOMODCACHE=D:\GoPathRepository\pkg\mod
go env -w GOROOT=D:\softwares\golang18
go env -w GOBIN= (设为空值)
go env -w GO111MODULE=auto (或SETX GO111MODULE auto)
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/ (或 set GOPROXY=https://goproxy.cn,direct)

# VScode工具
Ctrl+Shift+P 输入Go:Install/Update Tools
配置settings.json (见vscode-settings.json)
```

3. 通用编译&执行
```
# win编译linux可执行文件
set CGO_ENABLED=0
set GOOS=linux 
set GOARCH=amd64 
go build

# linux编译win可执行文件
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

# win运行
编译后运行$ ./moduledemo.exe
直接运行$ go run main.go
```

3. go mod模式配置
```
# 初始化模块
go mod init moduledemo

# 拉取依赖
go mod tidy

# 编译
go build
go build xx.go yy.go
```

4. go path模式配置
```
# 如果用gopath要建立三个目录
bin：存可执行文件
pkg：存编译的中间文件
src：存代码

# 仅下载依赖
go get 依赖名

# 下载及安装依赖
go install 依赖名

# 编译
go build
go build xxPackage
```