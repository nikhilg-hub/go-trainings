package main

// simple demo
// create goroutine on the fly and call it.
// goroutine can't return a value

import (
	"fmt"
	"time"
)

func routine(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "from routine:", name)
		time.Sleep(100 * time.Millisecond)
	}
}

func wrong() int {
	return 100
}

func main() {

	// spawn a goroutine:

	routine("routine 1")

	routine("main")

	// go func() {
	// 	i := 1
	// 	for {
	// 		fmt.Println(i)
	// 		time.Sleep(time.Millisecond * 100)
	// 		i++
	// 	}
	// }()

	// time.Sleep(time.Second)

}
