package main

import (
	util "exercises/util"
	"os"
)


func main() {
    searchString := os.Args[1]
    paths := util.GetPaths(true, true)

    util.GrepFiles(paths, searchString)
}