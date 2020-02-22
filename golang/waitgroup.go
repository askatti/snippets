package main

import "fmt"
import "time"
import "sync"

func main() {
        var TestCount int = 4
        var wg sync.WaitGroup
        wg.Add(TestCount)

        for i := 0; i < TestCount; i++ {

                go func( int) {
                        testcase(i)
                        wg.Done()
						fmt.Printf("Wg.Done called\n")
                }(i)
        }
        wg.Wait()
		fmt.Printf("TestEnd")
}

func testcase(cnt int) {
        for i := 1; i < cnt; i++ {
                fmt.Printf("Test Case Num=%v and counting till:%v\n",cnt, i)
                time.Sleep(1 * time.Second)
        }
}

