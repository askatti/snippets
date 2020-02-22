package main

import (
    "fmt"
    "os"
	"errors"
)

func main() {
    //if err := IsFileExists("/home/askatti/snippets-program"); err != nil {
    if err := IsDir("/home/askatti/snippets-program/"); err != nil {
        fmt.Println("Error:",err)
    } else {
        fmt.Println("Example file exists")
    }
	arrstr := []string{"/yome/askati/asdfdasf:/asdfads/asdfadsf/","/home/oo/:/home/kkk/","/home/ggg/:/home/eeee"}
    fmt.Println("arrstr[0]",arrstr[0])
    fmt.Println("arrstr[1:]",arrstr[1:])
}

// IsFileExists returns whether the given file exists or not and is (not a dir)
func IsFileExists(filename string) error {
     info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return err
    }
    if (info.IsDir()) {
        return errors.New("Its not a file")
    }
    return nil
}

// IsDir returns whether the given dir exists or not and is (not a file)
func IsDir(dirname string) error {
     info, err := os.Stat(dirname)
    if os.IsNotExist(err) {
        return err
    }
    if (info.IsDir()) {
        return nil
    }
     return errors.New("Its not a Dir")
}




