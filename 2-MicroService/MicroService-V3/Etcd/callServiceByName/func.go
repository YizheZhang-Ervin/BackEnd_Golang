package util

import (
	"context"
	"net/http"
)

//request是调用参数，response是返回结果
type Endpoint func(ctx context.Context, requestParam interface{}) (responseResult interface{}, err error)

//决定请求path以及参数
type EncodeRequestFunc func(context.Context, *http.Request, interface{}) error
