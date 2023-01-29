package utils

type MyError struct {
	Code    int
	Message string
}

func NewMyError(code int, msg string) error {
	return &MyError{Code: code, Message: msg}
}

func (this *MyError) Error() string {
	return this.Message
}
