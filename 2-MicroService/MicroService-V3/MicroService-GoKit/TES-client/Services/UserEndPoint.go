package Services

type UserRequest struct { //封装User请求结构体
	Uid    int `json:"uid"`
	Method string
}

type UserResponse struct {
	Result string `json:"result"`
}
