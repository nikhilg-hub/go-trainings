package main

import (
	"context"
	"fmt"
	"time"
)

// singnaling goroutine using channel
// signaling using Context
// withTimeout and withDeadline

func main() {

	ctxt := context.Background()

	// ctxt, cancelFunc := context.WithCancel(ctxt)

	// ctxt, cancelFunc := context.WithTimeout(ctxt, time.Second*1)

	ctxt, cancelFunc := context.WithDeadline(ctxt, time.Now().Add(time.Second))

	go PrintInfinite(ctxt)

	time.Sleep(2 * time.Second)

	// stop <- struct{}{}
	cancelFunc()

	time.Sleep(2 * time.Second)

	fmt.Println("Exiting from main")

}

func PrintInfinite(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Returning from PrintInfinite...", ctx.Err().Error())
			return

		default:
			fmt.Println(i)
			i++
		}
	}
}
