package main

import "fmt"
import "strings"
//import "encoding/json"
//import "bytes"

func main() {

var b = []string{}
var x = []string{}

for ponport :=1; ponport< 4;ponport++ {
 
	//t := fmt.Sprintf("{\"onu_info\": {\"info\": [{\"pon_port_id\": %d },]}}", ponport)
	a := fmt.Sprintf("{\"pon_port_id\": %d },", ponport)
	b = append(b,a)

	}
	temp := strings.TrimSuffix(string(b[len(b) -1]),",")
	x = b[0:len(b) -1]
	x = append(x, temp)
	//fmt.Println("temp=",temp)
	t := fmt.Sprintf("{\"onu_info\": {\"info\": %s}}", x)
		fmt.Println(t)
	}

