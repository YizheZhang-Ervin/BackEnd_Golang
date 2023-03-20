package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	mygomicroclient "my-gomicro-client/proto/my-gomicro-server"

	"github.com/micro/go-micro/v2/client"
)

func MyGomicroClientCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	// 获取数据，解析，存map
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	// 初始化客户端，调用远程服务
	mygomicroclientClient := mygomicroclient.NewMyGomicroClientService("go.micro.service.mygomicroclient", client.DefaultClient)
	rsp, err := mygomicroclientClient.Call(context.TODO(), &mygomicroclient.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	// 响应数据封装成map格式存储
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	// 封装的map编码发给浏览器
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
