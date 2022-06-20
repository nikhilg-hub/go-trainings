package main

import (
	"fmt"
	"sync"
	"time"
)

// race-condition
// solution using :
// channel, sync.Mutex, sync/atomic

var number = 0 // 1K

var lock = new(sync.Mutex)

func main() {

	// ch := make(chan struct{}, 1)

	for i := 0; i < 1000; i++ {
		go increment(lock)
	}

	time.Sleep(time.Second * 5) // wait for some time.

	fmt.Println(number)
}

func increment(mutex *sync.Mutex) {

	// ch <- struct{}{}
	mutex.Lock() // aquire

	number++

	mutex.Unlock() // release

}
