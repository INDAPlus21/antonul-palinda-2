package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	wg.Add(1)
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	wg.Done()
}
