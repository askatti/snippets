package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	//ch <- true
	var done sync.WaitGroup
	done.Add(1)
	go func() {
	Loop:
		for i := 0; i < 10; i++ {
			fmt.Println("before select..:")
			select {
			case val, ok := <-ch:
				if val == "" && ok == false {
					fmt.Println("Channel Closed")
					break Loop
				} else {
					fmt.Println("Channel Open")
				}
				//fmt.Println("Thread: is Channel closed")
			default:
				if i == 6 {
					//ch <- "done"
					close(ch)
				}
				fmt.Println("waiting:", i)
			}
		}

		fmt.Println("after select..:")
		//time.Sleep(5 * time.Second)

		fmt.Println("for loop oddutside")
		done.Done()
	}()
	done.Wait()
	//state := <-ch
	//fmt.Println(state)
	//time.Sleep(5 * time.Second)
	//close(ch)
	//time.Sleep(20 * time.Second)
	fmt.Println("Main done")
}
