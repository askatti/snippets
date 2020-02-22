package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    matches, _ := filepath.Glob("/tmp/sample_manifest/manifest1.yml")
    fmt.Println(matches)
    for _, match := range matches {
        fmt.Println(match)
    }
}
