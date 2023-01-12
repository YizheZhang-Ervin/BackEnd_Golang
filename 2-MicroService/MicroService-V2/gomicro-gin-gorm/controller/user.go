package controller

import (
	"github.com/gin-gonic/gin"
	"bj38web/web/utils"
	"net/http"
	"fmt"
	getCaptcha "bj38web/web/proto/getCaptcha" // 给包起别名
	"image/png"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro"
	"context"
	"encoding/json"
	"github.com/afocus/captcha"

	userMicro "bj38web/web/proto/user" // 给包起别名
)

// 获取 session 信息.
func GetSession(ctx *gin.Context) {
	// 初始化错误返回的 map
	resp := make(map[string]string)

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}

// 获取图片信息
func GetImageCd(ctx *gin.Context) {
	// 获取图片验证码 uuid
	uuid := ctx.Param("uuid")

	// 指定 consul 服务发现
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getCaptcha", consulService.Client())

	// 调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("未找到远程服务...")
		return
	}

	// 将得到的数据,反序列化,得到图片数据
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	// 将图片写出到 浏览器.
	png.Encode(ctx.Writer, img)

	fmt.Println("uuid = ", uuid)
}

// 获取短信验证码
func GetSmscd(ctx *gin.Context) {
	// 获取短信验证码
	phone := ctx.Param("phone")
	// 拆分 GET 请求中 的 URL === 格式: 资源路径?k=v&k=v&k=v
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	// 指定Consul 服务发现
	consulReg := consul.NewRegistry()

	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	// 初始化客户端
	microClient := userMicro.NewUserService("go.micro.srv.user", consulService.Client())

	// 调用远程函数:
	resp, err := microClient.SendSms(context.TODO(), &userMicro.Request{Phone: phone, ImgCode: imgCode, Uuid: uuid})
	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		return
	}

	// 发送校验结果 给 浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 发送注册信息
func PostRet(ctx *gin.Context) {
	/*	mobile := ctx.PostForm("mobile")
		pwd := ctx.PostForm("password")
		sms_code := ctx.PostForm("sms_code")

		fmt.Println("m = ", mobile, "pwd = ", pwd, "sms_code = ",sms_code)*/
	// 获取数据
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}

	ctx.Bind(&regData)

	fmt.Println("获取到的数据为:", regData)
}
