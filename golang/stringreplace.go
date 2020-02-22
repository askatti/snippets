package main

import (
    "fmt"
    "strings"
)

func main() {

/*
    // Example 1: Willkommen to GoLangCode.com
    myText := "Welcome to GoLangCode.com"
    myText = strings.Replace(myText, "Welcome", "Willkommen", -1)
    fmt.Println(myText)
*/
    // Example 2: Change first occurance
    // Output: The car sounds sound
    myText := []string{"/home/asdfas/asdf.conf:/afads/asdfdg/tgllg.yml","/homer/asdfads/ABC:/home/asdfads/EFG:rw","/home/assa/ss:/home/ss/hhh/"}
	k1 := strings.Split(myText[0],".")
	if len(k1) > 1 {
	   fmt.Println("contains dot in file",k1[1])
    } else {
       fmt.Println("No dot in file",k1[0])
      }
    fmt.Printf("split text=%+v",k1)
    o1 := myText[0]
    //a1 :=strings.Split(o1,":")
    //a1 = strings.Join(a1,"")
	//fmt.Println("split string=",a1)
	fmt.Println("current text",myText)
    a1 := strings.Replace(o1, "/homer/asdfads/ABC", "/ggg/ggg/ggg", 1)
	fmt.Println("string replace output=",a1)
	myText[0] = a1 
    fmt.Println("Updated Text",myText[0])
    fmt.Println("Updated Text",myText[1])
/*
    // Example 3: Replacing quotes (double backslash needed)
    // Output: I \'quote\' this text
    myText = "I 'quote' this text"
    myText = strings.Replace(myText, "'", "\\'", -1)
    fmt.Println(myText)
*/
}
