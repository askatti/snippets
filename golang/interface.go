package main
import "fmt"

func contains(k []interface{}, b interface{}) bool{
for _, a := range k {
		if a == b {
			fmt.Println("True")
			return true
		}
	}
    fmt.Println("False")
	return false
}

func main() {
	var v1 []interface{} = []int{1,2,3,4}
	var v2 interface{} = "Hi"
	contains(v1,v2)
}
