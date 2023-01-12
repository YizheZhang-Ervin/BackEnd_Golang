package handler

import (
	"context"

	getCaptcha "bj38web/service/getCaptcha/proto/getCaptcha"
	"github.com/afocus/captcha"
	"image/color"
	"encoding/json"
	"bj38web/service/getCaptcha/model"
)

type GetCaptcha struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *getCaptcha.Request, rsp *getCaptcha.Response) error {
	// 生成图片验证码

	// 初始化对象
	cap := captcha.New()

	// 设置字体
	cap.SetFont("./conf/comic.ttf")

	// 设置验证码大小
	cap.SetSize(128, 64)

	// 设置干扰强度
	cap.SetDisturbance(captcha.NORMAL)

	// 设置前景色
	cap.SetFrontColor(color.RGBA{0,0,0, 255})

	// 设置背景色
	cap.SetBkgColor(color.RGBA{100,0,255, 255}, color.RGBA{255,0,127, 255}, color.RGBA{255,255,10, 255})

	// 生成字体
	img, str := cap.Create(4, captcha.NUM)

	// 存储图片验证码到 redis 中
	err := model.SaveImgCode(str, req.Uuid)
	if err != nil {
		return err
	}

	// 将 生成成的图片 序列化.
	imgBuf, _ := json.Marshal(img)

	// 将 imgBuf 使用 参数 rsp 传出
	rsp.Img = imgBuf

	return nil
}

