package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
)

const CHUNK_SIZE = 64000

// Error handling
func getError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}

// Utils functions

func get_paths(start string) []string {

	files := make([]string, 0)

	PermissionDenied := errors.New("permission denied")

	filepath.Walk(start, func(path string, info fs.FileInfo, err error) error {

		getError(err)

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
				getError(err)
			}

		} else {
			getError(err)
		}

		defer f.Close()

		return nil
	})

	return files
}

func comp_files(file1 string, file2 string) bool {

	fi1, err := os.Stat(file1)

	getError(err)

	fi2, err := os.Stat(file2)

	getError(err)

	if fi1.Size() != fi2.Size() || fi1.Size() == 0 { // if they don't have the same size, they can't be the same file.
		return false
	} else {

		fi1, err := os.Open(file1)

		getError(err)

		fi2, err := os.Open(file2)

		getError(err)

		for {
			cb1 := make([]byte, CHUNK_SIZE)
			_, err1 := fi1.Read(cb1)

			getError(err1)

			cb2 := make([]byte, CHUNK_SIZE)
			_, err2 := fi2.Read(cb2)

			getError(err2)

			return bytes.Equal(cb1, cb2)

		}

	}

}

const MAX_CONCURRENT_JOBS = 10

func main() {

	files := get_paths(os.Args[1])

	barsize := len(files) * len(files)
	bar := progressbar.Default(int64(barsize))
	var equal_results string
	count := 0
	guard := make(chan bool, MAX_CONCURRENT_JOBS)

	for _i := 0; _i < len(files); _i++ {
		for _j := 0; _j < len(files); _j++ {

			guard <- true

			go func(i int, j int) {

				defer func() {
					<-guard
				}()

				if comp_files(files[i], files[j]) && i != j {
					result := fmt.Sprintf("EQUAL \n%v \n%v \n\n", files[i], files[j])
					equal_results += result
					count++
				}

			}(_i, _j)

			bar.Add(1)

		}
	}

	for i := 0; i < MAX_CONCURRENT_JOBS; i++ {
		guard <- true
	}

	fmt.Print(equal_results)

	fmt.Printf("You Have %v Duplicated Files!", count)

}
