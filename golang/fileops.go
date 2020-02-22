package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/home/askatti/my-go-workspace/src/controller/examples/sample_manifest.tar.gz"
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("File %s: Exists.\n", path)
	} else if os.IsNotExist(err) {
		fmt.Printf("'File %s' - Doesn't Exist.\n\tError=%s\n", path, err)
	} else {
		fmt.Printf("File status unknown \n")
	}
}
