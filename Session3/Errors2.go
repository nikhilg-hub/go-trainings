package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("Divide by zero error") // new error object

func Divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, DivideByZeroError
		return 0, fmt.Errorf("Cannot divide %d by %d", a, b)
	}
	return a / b, nil
}

func main() {

	result, err := Divide(12, 0)

	if err != nil {
		// error handling
		// fmt.Println("error:", err.Error())

		switch {
		case errors.Is(err, DivideByZeroError):
			fmt.Println("It's divide by zero error")
		default:
			fmt.Println("some other error: ", err.Error())
		}

		return
	}

	fmt.Println("result=", result)

}
