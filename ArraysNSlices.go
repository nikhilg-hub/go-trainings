package main

import "fmt"

/*

type Slice struct {
	ptr  // pointer to the first element of slice
	len // lenght of a slice
	cap // capacity of underlying array
}

*/

func main() {

	arr := []int{1, 2, 4, 5, 6, 7, 8} // try specifying size to demo an array
	// 0  1	 2  3  4  5  6

	// fmt.Printf("%T\n", arr)

	slice := arr[1:6]

	// fmt.Printf("%T\n", slice)

	PrintInfo(arr)
	PrintInfo(slice)

	slice[0] = 100

	fmt.Println("After Change")

	PrintInfo(arr)
	PrintInfo(slice)

	slice = append(slice, 200)
	slice = append(slice, 300) // 8

	slice[0] = 500

	fmt.Println("After append")

	PrintInfo(arr)
	PrintInfo(slice)

	// Ranging over a slice

	for i, _ := range slice {
		fmt.Println(i, "->")
	}

}

func PrintInfo(s []int) {
	fmt.Printf("data = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
