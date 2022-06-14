package main

import "fmt"

// func [(v Type)] identifier(identifier datatype) (returntype, returntype2) {

// }

func main() {
	// fmt.Println(Sum(12, 13))
	// a, b := Bisect(13)
	// fmt.Println(a, b)

	fmt.Println(SumAll(1, 2, 3, 4, 5))

}

func Sum(i, j int) int {
	return i + j
}

func Bisect(i int) (k, j int) {
	if i%2 == 0 {
		k, j = i/2, i/2
	} else {
		k, j = i/2, (i/2)+1
	}
	return
}

func SumAll(arr ...int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
