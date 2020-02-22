package main

// Here we show how to properly terminate multiple go routines by using a context.
// Thanks to WaitGroup we'll be able to end all go routines gracefully before the main function ends.

import (
	"fmt"
	"os"
	"context"
	"sync"
	"math/rand"
	"time"
)

// it prints all values pushed into "ch" ("ch" here is read only)
func reader(wg *sync.WaitGroup, ctx context.Context, ch <-chan int) {
	wg.Add(1) // adds delta, if the counter becomes zero, all goroutines blocked on Wait are released
	defer wg.Done() // decrements the WaitGroup counter by one when the function returns

	for {
		select {
		case <-ctx.Done(): // Done returns a channel that's closed when work done on behalf of this context is canceled
			fmt.Println("Exiting from reading go routine")
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel has been closed")
				return
			}

			fmt.Println(v)
		}
	}
}

// it writes a random integer into "ch" every second ("ch" here is write only)
func writer(wg *sync.WaitGroup, ctx context.Context, ch chan<- int) {
	wg.Add(1) // adds delta, if the counter becomes zero, all goroutines blocked on Wait are released
	defer wg.Done() // decrements the WaitGroup counter by one when the function returns

	for {
		select {
		case <-ctx.Done(): // Done returns a channel that's closed when work done on behalf of this context is canceled
			fmt.Println("Exiting from writing go routine")
			return
		default:
			ch<- rand.Intn(100) // pushes a random integer from 0 to 100 into the channel
			time.Sleep(1 * time.Second) // sleeps one second
		}
	}
}

func main() {
	channel := make(chan int) // unbuffered channel, could use a buffered one too
	waitGroup := sync.WaitGroup{} // a WaitGroup waits for a collection of goroutines to finish, pass this by address
	// context.WithCancel returns a copy of parent with a new Done channel.
	// The returned context's Done channel is closed when the returned cancel function is called or when the parent
	// context's Done channel is closed, whichever happens first.
	ctx, cancel := context.WithCancel(context.Background())

	go reader(&waitGroup, ctx, channel) // go routine that prints all values pushed into "channel"
	go writer(&waitGroup, ctx, channel) // go routine that writes a random integer into "channel" every second

	// go routine that listens for an Enter keystroke to terminate the program
	go func() {
		os.Stdin.Read(make([]byte, 1)) // wait for Enter keystroke
		cancel() // cancel the associated context
	}()

	waitGroup.Wait() // it blocks until the WaitGroup counter is zero
}
