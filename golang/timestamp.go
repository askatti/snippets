package main

import (
	"fmt"
	"time"
	)

func main () {

	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(time.Now())
	fmt.Println(time.Now().String())
	t1 := time.Now().UnixNano()/1000000
	fmt.Println(t1)
	t2 := time.Unix(t1,0)
	fmt.Println(t2)
}
