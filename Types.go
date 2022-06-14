package main

import "fmt"

func main() {

	// String1 := "First String"
	// String2 := `String
	// With
	// New Line`
	// var str string

	// fmt.Printf("%q", str)

	// fmt.Println(String1)
	// fmt.Println(String2)

	// var boolean bool = true
	// fmt.Println(boolean)

	// i, i32, i64 := 1, int32(2), int64(3)

	// fmt.Printf("type of i = %T\ntype of i32 = %T\ntype of i64 = %T", i, i32, i64)

	// f := 90
	// fmt.Printf("%v is of type %T\n", f, f)

	// cmplx := 34 - 45i
	// fmt.Printf("%v is of type %T\n", cmplx, cmplx)

	// fmt.Println([]rune("golang"))

	for i, v := range "golang" {
		fmt.Println(i, "->", v)
	}

}
