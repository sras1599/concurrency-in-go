package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	util "exercises/util"
)

func catFile(path string, wg *sync.WaitGroup) {
    defer wg.Done()

    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        fmt.Println(err)
        return
    }

    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        fmt.Println(err)
        return
    }

    fileName := strings.Split(path, "/")[len(strings.Split(path, "/")) - 1]

    output := fmt.Sprintf("Contents of %s\n\n", fileName)
    output += string(bs)
    output += "----------------"

    println(output)
}

func catFiles(paths []string) {
    var wg sync.WaitGroup

    for _, path := range paths {
        wg.Add(1)
        go catFile(path, &wg)
    }

    wg.Wait()
}

func main() {
    paths := util.GetPaths(false, false)

    catFiles(paths)
}