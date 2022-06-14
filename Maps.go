package main

import "fmt"

func main() {

	// first way:

	// map1 := map[int]string{
	// 	1: "Ek",
	// 	2: "Do",
	// 	3: "Teen",
	// }

	map2 := make(map[int]string)
	map2[4] = "Chaar"
	map2[5] = "Panch" // delete
	map2[6] = "Chhe"

	// Access a value using key:
	// fmt.Println("map1[3] =", map1[3])
	// fmt.Println("map1[4] =", map1[4]) // no key fouund case (will return zero value for a type)

	// Check if a key is present or not:
	// if value, isPresent := map2[5]; isPresent {
	// 	fmt.Println(value)
	// }

	// delete a key value pair:
	delete(map2, 5)

	// if value, isPresent := map2[5]; isPresent {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("No key found")
	// }

	// ranging over a map :
	for k, v := range map2 {
		fmt.Println(k, ":", v)
	}

	// Note: different runs may provide defferent results of ranging over a map.
}
