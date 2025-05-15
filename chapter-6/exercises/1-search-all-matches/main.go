package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

type matches []string

func searchDir(path string, filename string, wg *sync.WaitGroup, matches *matches) {
	files, _ := os.ReadDir(path)

	for _, file := range files {
		fpath := filepath.Join(path, file.Name())

		if strings.Contains(file.Name(), filename) {
			*matches = append(*matches, fpath)
		}

		if file.IsDir() {
			wg.Add(1)
			go searchDir(fpath, filename, wg, matches)
		}
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	fileMatches := make(matches, 0)
	slices.Sort(fileMatches)

	wg.Add(1)
	go searchDir(os.Args[1], os.Args[2], &wg, &fileMatches)
	wg.Wait()

	if len(fileMatches) > 0 {
		fmt.Printf("Files with the term `%s` in them was found in the following paths:\n", os.Args[2])

		for _, path := range fileMatches {
			fmt.Println(path)
		}
	}
}
