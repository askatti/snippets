package main

import (
    "fmt"
    //"os"
)

func main() {

    serverport := "localhost:6379"
//    fmt.Println("\nserverpo: %s \n",serverport)
    dbinit(serverport)
	/*/
	err := Set("MANIFEST","")
	if err != nil {
            fmt.Println("\nerr2: Error in GET:",err)
    }
	*/
	/*
	data, err2 := Get("MANIFEST")
     if err2 != nil {
            fmt.Println("\nerr2: Error in GET:",err2)
    }
	if data == "" {
	fmt.Printf("\ndata: GET MANIFEST:>%s<",data)
	}

    err1 := Delete("MANIFEST")
        if err1 != nil {
            fmt.Printf("\nerr1: DEL Error:>%s<",err1)
    }
	/*
	err3 := Set("MANIFEST","")
     if err3 != nil {
            fmt.Println("\nerr2: Error in SET:",err3)
    }
	*/
	data1, err4 := Subscribe("MANIFEST_BLOCK1")
     if err4 != nil {
            fmt.Println("\nerr4 Error in Subscribe:",err4)
    }
	if data1 == "" {
	fmt.Printf("\ndata1: 4. GET MANIFEST_BLOCK Empty return: output:>%s<",data1)
	} else {
	fmt.Printf("\ndata1: 5. GET MANIFEST_BLOCK output:>%s<",data1)
	}
}
