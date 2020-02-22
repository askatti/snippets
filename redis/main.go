package main

import (
    "fmt"
    //"os"
)

func main() {

    serverport := "localhost:6379"
    fmt.Println("serverpo: %s ",serverport)
    dbinit(serverport)

    err1 := Set("MANIFEST_BLOCK1", "FLIP1")
    if err1 != nil {
       fmt.Println("SET Error",err1)
       return
    }
    data,err1 := Get("MANIFEST_BLOCK1")
    fmt.Println("\nGET for key NAME:=",data)

//   myvar, err2 := Subscribe("NAME")
//   fmt.Println("Return from subscribe for key NAME=: Err=\n",myvar, err2)
  // Publish("NAME", "0")
}
