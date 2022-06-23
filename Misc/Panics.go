package main

import "fmt"

const (
	A = (iota + 1) * 2 // 0
	B                  // 2
	C                  // 4
	D                  // 6
)

func main() {

	fmt.Println("Before Divide")

	fmt.Println(Divide(100, 0)) //

	fmt.Println("After Divide")

	// fmt.Println(A, B, C, D)

	// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// sl := arr[:5] // 5, 10

	// fmt.Println(sl)
	// fmt.Println("len :", len(sl))
	// fmt.Println("cap :", cap(sl))

	// ch := make(chan int)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i + 1
	// 	}
	// 	close(ch)
	// }()

	// for v := range ch {
	// 	fmt.Println(v)
	// }

	// mmm := make(map[int]string)
	// mmm[10] = "Ten"
	// mmm[20] = "Twenty"
	// mmm[30] = "Thirty"

	// fmt.Println(mmm)

	// delete(mmm, 10)

	// fmt.Println("after delete:", mmm)

	// dummy := new(Dummy)
	// dummy.Id = 10
	// dummy.Name = "ABC"
	// fmt.Println(dummy)

	// fmt.Println("Before panic")

	// PanicHere()

	// fmt.Println("After panic")

}

// func PanicHere() {
// 	panic("Panic in main routine")
// }

type Dummy struct {
	Id   int
	Name string
}

func Divide(a, b int) int {

	defer DeferredFunc()

	return a / b // panic

}

func DeferredFunc() {
	fmt.Println("This is deferred function call and you can revover here.", recover())
}
