package main

import (
         "os"
         "github.com/opencord/cordctl/pkg/format"
	"fmt"
	"reflect"
)

func main() {
        type K struct {
              C int
              D string
	}
	type T struct {
		A int
		B string
		G []K
	}
        
        l1 := []K{{11,"mallu"}, {321,"kallu"}}

	t := T{23, "skidoo",nil}
	
	t.G = l1

	
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	tableFmt := format.Format("table{{.A}}\t{{.G}}\t")
	if err := tableFmt.Execute(os.Stdout, true, t); err != nil {
	fmt.Printf("Error.............")
	}
}

