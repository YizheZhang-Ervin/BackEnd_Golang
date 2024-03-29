// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `json:"name,options=you|me"`
	Gender string `json:"gender"`
}

type Response struct {
	Message string `json:"message"`
	Data any `json:"data"`
}

type IdRequest struct {
	Id string `json:"name" path:"id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}