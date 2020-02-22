package main

import (
  "fmt"
)

type Person struct {
  age int
}

func main() {
  x := Person{}

  switch {
    case x == Person{1}:
        fmt.Println("Structure is empty")
    default:
        fmt.Println("Structure is not empty")
  }
}
