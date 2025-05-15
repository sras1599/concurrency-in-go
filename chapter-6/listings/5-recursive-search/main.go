package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func searchDir(path string, filename string, wg *sync.WaitGroup) {
	files, _ := os.ReadDir(path)

	for _, file := range files {
		fpath := filepath.Join(path, file.Name())

		if strings.Contains(file.Name(), filename) {
			fmt.Printf("File containing the term `%s` found in: %s\n", filename, fpath)
		}

		if file.IsDir() {
			wg.Add(1)
			go searchDir(fpath, filename, wg)
		}
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go searchDir(os.Args[1], os.Args[2], &wg)
	wg.Wait()
}
