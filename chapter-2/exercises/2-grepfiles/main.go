package main

import (
	util "exercises/util"
	"os"
)


func main() {
    searchString := os.Args[1]
    paths := util.GetPaths(true, false)

    util.GrepFiles(paths, searchString)
}