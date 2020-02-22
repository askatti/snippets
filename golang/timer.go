package main
import (
	"time"
	"fmt"
)

func main() {

ticker := time.NewTicker(10 * time.Second)
defer ticker.Stop()
fmt.Printf("10 sec timer expir")

}
