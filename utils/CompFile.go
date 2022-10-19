package utils

import (
	"bytes"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}

const CHUNK_SIZE = 64000

func GetPaths(start string) []string {

	files := make([]string, 0)

	PermissionDenied := errors.New("permission denied")

	filepath.Walk(start, func(path string, info fs.FileInfo, err error) error {

		GetError(err)

		f, err1 := os.Open(path)

		if err1 != PermissionDenied {

			finfo, err := f.Stat()

			if err == nil {
				if !(finfo.IsDir()) { // avoid directories

					if strings.Contains(path, ".") {

						files = append(files, path)

					}
				}

			} else {
				GetError(err)
			}

		} else {
			GetError(err)
		}

		defer f.Close()

		return nil
	})

	return files
}

func CompFiles(file1 string, file2 string) bool {

	fi1, err := os.Stat(file1)

	GetError(err)

	fi2, err := os.Stat(file2)

	GetError(err)

	if fi1.Size() != fi2.Size() || fi1.Size() == 0 { // if they don't have the same size, they can't be the same file.
		return false
	} else {

		fi1, err := os.Open(file1)

		GetError(err)

		fi2, err := os.Open(file2)

		GetError(err)

		for {
			cb1 := make([]byte, CHUNK_SIZE)
			_, err1 := fi1.Read(cb1)

			GetError(err1)

			cb2 := make([]byte, CHUNK_SIZE)
			_, err2 := fi2.Read(cb2)

			GetError(err2)

			return bytes.Equal(cb1, cb2)

		}

	}

}
