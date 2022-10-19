package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	cf "github.com/Shinji-Mimura/gopare/utils"
	"github.com/schollz/progressbar/v3"
)

func main() {

	if len(os.Args) != 3 {
		log.Fatal("Missing or Excessive Arguments!\nUse: ./gopare <DIRECTORY> <THREADS NUMBER>\nRecommended Threads Number = 10")
		os.Exit(0)
	}

	MAX_CONCURRENT_JOBS, _ := strconv.Atoi(os.Args[2])

	directory := os.Args[1]

	files := cf.GetPaths(directory)

	barsize := (len(files) * len(files))

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

				if cf.CompFiles(files[i], files[j]) && i != j {
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

	fmt.Printf("You Have %v Duplicated Files!\n", count)

}
