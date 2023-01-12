package main

import (
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main()  {
	// 初始化对象
	cap := captcha.New()

	// 设置字体
	cap.SetFont("comic.ttf")

	// 设置验证码大小
	cap.SetSize(128, 64)

	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)

	// 设置前景色
	cap.SetFrontColor(color.RGBA{0,0,0, 255})

	// 设置背景色
	cap.SetBkgColor(color.RGBA{100,0,255, 255}, color.RGBA{255,0,127, 255}, color.RGBA{255,255,10, 255})

	// 生成字体 -- 将图片验证码, 展示到页面中.
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img, str := cap.Create(4, captcha.NUM)
		png.Encode(w, img)

		println(str)
	})

	// 启动服务
	http.ListenAndServe(":8086", nil)
}
