package main

import "fmt"

// channel basics : how to make, len parameter, write to a channel, read from a channel, range over a channel.

// closing a channel

// read/write Blocking nature of a channel.

// send only and receive only channels.

// producer consumer using channels.

func main() {

	ch := GetChannel()

	for v := range ch {
		fmt.Println("consumed", v)
	}

}

func Cons(ch chan int) {
	for v := range ch {
		fmt.Println("consumed", v)
	}
}

func Prod(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		fmt.Println("produced", i+1)
	}

	close(ch)
}

func GetChannel() chan int {

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i + 1
			fmt.Println("produced", i+1)
		}
		close(ch)
	}()

	return ch
}
