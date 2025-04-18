package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

func isDir(path string) bool {
    info, err := os.Stat(path)

    if err != nil {
        fmt.Println(err)
        return false
    }

    return info.IsDir()
}

func GetPaths(grep bool, recursive bool) []string {
    minArgs := 2
    if grep {
        minArgs = 3
    }
    
    if len(os.Args) < minArgs {
        fmt.Println("Please provide a space separated list of filenames.")
        os.Exit(1)
    }

    if grep {
        return removeDuplicates(processPaths(os.Args[2:], recursive))
    }
    
    return removeDuplicates(os.Args[1:])
}

func removeDuplicates(slice []string) []string {
    slices.Sort(slice)

    return slices.Compact(slice)
}

func processPath(path string, recursive bool) []string {
    var paths []string

    if isDir(path) {
        paths = append(paths, flattenPath(path, recursive)...)
    } else {
        absPath, _ := filepath.Abs(path)
        paths = append(paths, absPath)
    }

    return paths
}

func flattenPath(dirPath string, recursive bool) []string {
    var flattenedPaths []string
    absPath, _ := filepath.Abs(dirPath)
    
    
    paths, err := os.ReadDir(dirPath)
    if err != nil {
        log.Fatal(err)
    }

    for _, path := range paths {
        fullPath := filepath.Join(absPath, path.Name())
        
        if isDir(fullPath) {
            if !recursive {
                continue
            } else {
                flattenedPaths = append(flattenedPaths, flattenPath(fullPath, recursive)...)
            }
        } else {
            flattenedPaths = append(flattenedPaths, fullPath)
        }
    }

    return flattenedPaths
}

func processPaths(paths []string, recursive bool) []string {
    var ret []string
    
    for _, path := range paths {
        ret = append(ret, processPath(path, recursive)...)
    }

    return ret
}

func searchFile(path string, searchString string) bool {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
        return false
    }

    defer file.Close()

    bs := make([]byte, 1024)
    _, err = file.Read(bs)
    if err != nil {
        fmt.Println(err)
        return false
    }

    return strings.Contains(string(bs), searchString)
}

func grepFile(path string, searchString string, wg *sync.WaitGroup) {
    defer wg.Done()
    found := searchFile(path, searchString)
    
    execPath, _ := os.Getwd()
    cleanedPath, _ := filepath.Rel(execPath, path)

    output := fmt.Sprintf("Searching %s for %s", cleanedPath, searchString)
    if found {
        output += fmt.Sprintf("\nFound %s in %s", searchString, cleanedPath)
    } else {
        output += fmt.Sprintf("\nDid not find %s in %s", searchString, cleanedPath)
    }

    output += "\n-------------"
    fmt.Println(output)
}

func GrepFiles(paths []string, searchString string) {
    var wg sync.WaitGroup

    for _, path := range paths {
        wg.Add(1)
        go grepFile(path, searchString, &wg)
    }

    wg.Wait()
}