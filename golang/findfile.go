package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "/tmp/sample_manifest/"
	var matchfiles = []string{path + "topology.yml", path + "manifest.yml"}
	if  nomatch, _ := ValidateIsPresentManifestFiles(matchfiles); len(nomatch) > 0 {
		fmt.Println("main: mo match",nomatch)
	}else {
	  fmt.Println("Match found")
	}
}

// ValidateIsPresentManifestFiles return non matching files names string, or return nil
func ValidateIsPresentManifestFiles(path []string) (match string,err  error) {
	for _, mymatch := range path {
		nilmatch,_ := filepath.Glob(mymatch)
		if len(nilmatch) < 1{
			err = fmt.Errorf("No match found:",err)
			return mymatch,err
		}
	}
	return
}
