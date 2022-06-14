package main

import (
	"fmt"
	"sync"
	"time"
)

// Determine completion of all tasks using waitgroup.

func main() {

	// create a waitgroup
	wg := new(sync.WaitGroup) // remember always pointer should be used.

	// wg := sync.WaitGroup{} // this is wrong

	// data
	filenames := []string{"first", "second", "third", "fourth"}

	// adding number of tasks to waitgroup
	wg.Add(len(filenames)) // 4

	// spawn the goroutines:
	for _, fname := range filenames {

		// note the syntax go
		process(fname, wg)

	}

	// Bocks the flow until all tasks are completed.
	// wg.Wait() // becomes zero
}

func process(filename string, wg *sync.WaitGroup) {

	time.Sleep(2 * time.Second)

	fmt.Println("Processed:", filename)

	// Notify the waitgroup that one of the tasks is completed.
	wg.Done() // -1
}
