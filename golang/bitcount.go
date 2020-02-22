package main

import (
    //"math/bits"
    //"unsafe"
    "fmt"
)

func main() {
    value := uint(100)
    var val int;
    // Loop over all bits in the uint.
    for i := 0; i < 8; i++ {
        // If there is a bit set at this position, write a 1.
        // ... Otherwise write a 0.
        if value & (1 << uint(i)) != 0 {
            val++;
        } else {
            //fmt.Print("0")
        }
    }
    fmt.Println("Count=",val)
}
