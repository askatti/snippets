package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func WalkAllFilesInDir(dir string) error {
    return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
        if e != nil {
            return e
        }

        // check if it is a regular file (not dir)
        if info.Mode().IsRegular() {
            fmt.Println("file name:", info.Name())
            fmt.Println("file path:", path)
        }
        return nil
    })
}

func main() {
    WalkAllFilesInDir("/tmp/sample_manifest")
}
