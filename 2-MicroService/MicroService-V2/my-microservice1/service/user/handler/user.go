package handler

import (
	"context"
	"fmt"
	"math/rand"
	"my-microservice1/service/user/model"
	user "my-microservice1/service/user/proto/user"
	"my-microservice1/service/user/utils"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) SendSms(ctx context.Context, req *user.Request, rsp *user.Response) error {

	// 校验图片验证码 是否正确
	result := model.CheckImgCode(req.Uuid, req.ImgCode)
	if result {
		// 发送短信
		client, _ := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "xxx", "yyy")

		request := dysmsapi.CreateSendSmsRequest()
		request.Scheme = "https"

		request.Domain = "xx.xx.com" //域名  ---参考讲义补充!
		request.PhoneNumbers = req.Phone
		request.SignName = "xxweb"
		request.TemplateCode = "SMS_XXX"

		// 生成一个随机 6 位数, 做验证码
		rand.Seed(time.Now().UnixNano()) // 播种随机数种子.
		// 生成6位随机数.
		smsCode := fmt.Sprintf("%06d", rand.Int31n(1000000))

		request.TemplateParam = `{"code":"` + smsCode + `"}`

		response, _ := client.SendSms(request)
		if response.IsSuccess() {
			// 发送短信验证码 成功
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

			// 将 电话号:短信验证码 ,存入到 Redis 数据库
			err := model.SaveSmsCode(req.Phone, smsCode)
			if err != nil {
				fmt.Println("存储短信验证码到redis失败:", err)
				rsp.Errno = utils.RECODE_DBERR
				rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
			}

		} else {
			// 发送端验证码 失败.
			rsp.Errno = utils.RECODE_SMSERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		}

	} else {
		// 校验失败, 发送错误信息
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}
	return nil
}

func (e *User) Register(ctx context.Context, req *user.RegReq, rsp *user.Response) error {

	// 先校验短信验证码,是否正确. redis 中存储短信验证码.
	err := model.CheckSmsCode(req.Mobile, req.SmsCode)
	if err == nil {

		// 如果校验正确. 注册用户. 将数据写入到 MySQL数据库.
		err = model.RegisterUser(req.Mobile, req.Password)
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		} else {
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		}

	} else { // 短信验证码错误
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}

	return nil
}
