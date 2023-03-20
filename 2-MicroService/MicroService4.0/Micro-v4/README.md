# Go Micro v4

## 安装
```
go install github.com/go-micro/cli/cmd/go-micro@latest
```

## make (windows中)
```
下载https://osdn.dl.osdn.net/mingw/68260/mingw-get-setup.exe 或tdm-gcc
安装mingw32-make.exe，改名make.exe加入PATH
或直接D:\softwares\tdm-gcc\bin
```

## 创服务
```
go-micro new service helloworld
cd helloworld
make proto init update tidy
go-micro run
go-micro call helloworld Helloworld.Call '{"name": "John"}'
go-micro stream server helloworld Helloworld.ServerStream '{"count": 10}'
```

## 创函数
```
go-micro new function test-func
cd test-func
make init proto update tidy
```

## 创客户端
```
go-micro new client helloworld
creating client helloworld
cd helloworld-client
make tidy
go-micro run
```

## 描述服务
```
go-micro describe service helloworld
```