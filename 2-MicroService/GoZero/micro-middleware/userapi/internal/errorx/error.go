package errorx

var ParamsError = New(1101001, "参数不正确")

type BizError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int, msg string) *BizError {
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

func (e *BizError) Error() string {
	return e.Msg
}

func (e *BizError) Data() interface{} {
	return &ErrorResponse{
		e.Code,
		e.Msg,
	}
}
