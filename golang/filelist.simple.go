package main

import (
	"fmt"
	"os"
	"path/filepath"
	"errors"
)

func listfiles(rootpath string, files *[]string) error {
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		   if filepath.Ext(path) == ".yml"{
		     *files = append(*files, path)
		   }
		   return nil
	})
	if err != nil {
		fmt.Printf("File list Error=%s",err)
		return err
	}
    if len(*files) < 1{
    return errors.New("No files found in given path")
    }
	fmt.Println("Num of files found=",len(*files))
    return nil
}

func main() {
	root := "/home/askatti/my-go-workspace/src/controller/examples/sample_manifest/adapters/"
	var files []string

    if err := listfiles(root, &files); err != nil {
     fmt.Println("func returned error=", err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
