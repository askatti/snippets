package main

import (
	c "controller/cmd/common"
	"fmt"
	"strconv"
	"time"
)

func main() {

	serverport := "localhost:6379"
	c.Evdb.Dbinit(serverport, c.EVENTSTOREDB)

	var tclist = map[int]int64{1: 10, 2: 20, 3: 30}

	for {
		for id, tc := range tclist {
			time.Sleep(3 * time.Second)
			value2, _ := c.Evdb.Get("TC")
			temp, _ := strconv.Atoi(value2)
			if temp == 30 {
				fmt.Printf("TC=%d\n", value2)
				delete(tclist, temp)
			}
			fmt.Printf("Retrived stored key:TC value2=%s\n", value2)
			fmt.Printf("for loop id=%d tc=%d\n", id, tc)
		}
	}
}
