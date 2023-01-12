package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	test77 "test77/proto/test66"
)

func Test77Call(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json  --- 获取数据, 并解析, 存储至map 中.
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service  --- 初始化客户端, 并调用远程服务
	test77Client := test77.NewTest66Service("go.micro.srv.test66", client.DefaultClient)
	rsp, err := test77Client.Call(context.TODO(), &test77.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response  --- 将应答数据,封装程 map 格式存储.
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json  --- 将封装好的 map 编码, 并会发给 浏览器.
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
