package main

import (
	"fmt"
	"log"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer c.Close()

	ret, _ := c.Do("SET","fleet", "truck1", "POINT", "33", "-115")
	fmt.Printf("%v\n", ret)

	ret, _ = c.Do("GET","fleet", "truck1")
	fmt.Printf("%v\n", ret)

}
