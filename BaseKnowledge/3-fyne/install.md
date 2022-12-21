# fyne
## Links
```
# 官网
https://developer.fyne.io/started/

# 安装tdm-gcc
https://jmeubank.github.io/tdm-gcc/download/

# 升级msys2
https://www.msys2.org/
```

## 创建项目
```
go mod init myfyne
go get fyne.io/fyne/v2
go mod tidy
go run .
```

## 电脑端打包
```
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os darwin -icon myapp.png
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png
fyne install -icon myapp.png
```

## 移动端打包
```
# 安卓
fyne package -os android -appID com.example.myapp -icon mobileIcon.png
或adb install myapp.apk

# IOS
fyne package -os ios -appID com.example.myapp -icon mobileIcon.png
或xcrun simctl install booted myapp.app
```

## 发布
```
# macos
fyne release -appID com.example.myapp -appVersion 1.0 -appBuild 1 -category games

# android
fyne release -os android -appID com.example.myapp -appVersion 1.0 -appBuild 1

# ios
fyne release -os ios -appID com.example.myapp -appVersion 1.0 -appBuild 1
```

## Metadata
```
FyneApp.toml内容为：

Website = "https://example.com"
[Details]
Icon = "Icon.png"
Name = "My App"
ID = "com.example.app"
Version = "1.0.0"
Build = 1
```

## 给其他平台编译
```
go get github.com/fyne-io/fyne-cross
fyne-cross windows -arch=*
fyne-cross linux
fyne-cross linux -output bugs ./cmd/bugs
```

## 静态资源
```
# 生成
fyne bundle image.png >> bundled.go
fyne bundle -append image2.png >> bundled.go

# 加载
img := canvas.NewImageFromResource(resourceImagePng)
```