package main
import "fmt"
import "unsafe"

type SomeStruct struct {
	name string
}

func main() {
    v := SomeStruct{}
    s := getSomeStruct(v)
    if s.name == nil {
        fmt.Println("nil pointer")
        return
    } else {

        fmt.Println("not nil,printing name:",unsafe.Sizeof(s)) // It will crash here
    }
}

func getSomeStruct(s SomeStruct) *SomeStruct {
	return &s // This WILL compile
}
