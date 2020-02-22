package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func untar(inputFileName string, target string) ([]string, error) {

	var files []string
	var err error
	file, err := os.Open(inputFileName)
	if err != nil {
		return files, err
	}
	defer file.Close()

	var gzipfileReader io.ReadCloser = file
	// if the input file is in tar.gz format, handle it
	if strings.HasSuffix(inputFileName, ".gz") {
		if gzipfileReader, err = gzip.NewReader(file); err != nil {
			fmt.Println("Manifest is not a gzip file")
			return files, err
		}
		defer gzipfileReader.Close()
	}
	tarReader := tar.NewReader(gzipfileReader)

	// Extract tarred files
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return files, err
		}

		// get the individual filename and extract to the given dest directory
		filename := header.Name
		fpath := filepath.Join(target, filename)
		files = append(files, fpath)

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			err = os.MkdirAll(fpath, os.FileMode(header.Mode)) // using default Mode
			if err != nil {
				return files, err
			}

		case tar.TypeReg:
			// handle normal file
			writer, err := os.Create(fpath)

			if err != nil {
				return files, err
			}

			io.Copy(writer, tarReader)

			err = os.Chmod(fpath, os.FileMode(header.Mode))

			if err != nil {
				return files, err
			}

			writer.Close()
		default:
			return files, fmt.Errorf("Unable to untar type : %c in file %s", header.Typeflag, fpath)
		}
	}
	return files, nil
}
