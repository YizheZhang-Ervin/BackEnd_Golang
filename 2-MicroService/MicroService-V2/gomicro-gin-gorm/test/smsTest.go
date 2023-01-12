package main

import (
"fmt"
"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func main() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FgbQXjf117SX7E75Rmn", "6icOghQlhjevrTM5PxfiB8nDTxB9z6")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.Domain = "dysmsapi.aliyuncs.com"  //域名  ---参考讲义补充!
	request.PhoneNumbers = "18610382737"
	request.SignName = "爱家租房网"
	request.TemplateCode = "SMS_183242785"
	request.TemplateParam = `{"code":232323}`

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

