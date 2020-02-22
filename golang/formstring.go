package main

import "fmt"
//import "encoding/json"
//import "bytes"

func main() {
	
	oltserial := "hi"
	oltip := "byye"
	olthost := "which"
	oltmac := "amc"
	t := fmt.Sprintf("{\"serial_number\":\"%s\",\"ip\":\"%s\",\"hostname\":\"%s\",\"mac\":\"%s\"}", oltserial, oltip, olthost, oltmac)
	
	fmt.Println(t)
	}
