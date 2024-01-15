package Services

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func DecodeAccessRequest(c context.Context, r *http.Request) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	result := gjson.Parse(string(body)) //第三方库解析json
	if result.IsObject() {              //如果是json就返回true
		username := result.Get("username")
		userpass := result.Get("userpass")
		return AccessRequest{Username: username.String(), Userpass: userpass.String(), Method: r.Method}, nil
	}
	return nil, errors.New("参数错误")

}
func EncodeAccessResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response) //返回一个bool值判断response是否可以正确的转化为json，不能则抛出异常，返回给调用方
}
