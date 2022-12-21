package main

// go run interface.go
/*
	go build interface.go
	./basic
*/

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("--------------------------test 20 语言接口--------------------------")
	var phone Phone
	phone = new(IPhone)
	phone.call()

	fmt.Println("--------------------------test 21 错误处理--------------------------")
	// new实现error接口
	result, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// 手写error接口
	_, errorMsg := Divide(100, 0)
	if errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

// 接口
type Phone interface {
	call()
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("iPhone")
}

// 错误处理1
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	} else {
		return 0, errors.New("No Error")
	}
}

// 错误处理2
type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0 `
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}
