package main

import "fmt"

var n = 9

func main() {
	// r := 't'
	// fmt.Printf("%v, %T", r, r)

	i := 99

	fmt.Println("In main:", n)
}

func init() {
	n = 10
}

func init() {
	n = 11
}
