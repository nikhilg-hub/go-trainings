package main

import (
	"fmt"
)

// type DivideByZeroError struct {
// 	Int1, Int2 int
// }

// func (err *DivideByZeroError) Error() string {
// 	return fmt.Sprintf("DivideByZeroError : cannot divide %d by %d", err.Int1, err.Int2)
// }

// func main() {
// 	result, err := Divide(33, 0)
// 	if err != nil {
// 		var target *DivideByZeroError
// 		switch {
// 		case errors.As(err, &target):
// 			fmt.Println("Divide by zero error:", target)
// 		default:
// 			fmt.Println("Some other error in division:", err.Error())
// 		}

// 		return
// 	}
// 	fmt.Println(result)
// }

func main() {

	result, err := Divide(12, 6)

	if err != nil {
		// error handling
		fmt.Println("error:", err.Error())
		return
	}

	fmt.Println("result=", result)

}

// func Divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("Divide by zero error")
// 	}
// 	return a / b, nil
// }

// public void Produce() throws Exception{

// }
