package main

import (
	"fmt"
	"sync"
	"time"
)

// race-condition
// solution using :
// channel, sync.Mutex, sync/atomic

var number = 0

var lock = new(sync.Mutex)

func main() {

	for i := 0; i < 1000; i++ {
		go increment()
	}

	time.Sleep(time.Second * 5) // wait for some time.

	fmt.Println(number)
}

func increment() {
	number++
}
