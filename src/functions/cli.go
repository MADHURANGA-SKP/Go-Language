package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "the name to greet")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("hello, %s!\n", *name)
	} else if flag.Arg(0) == "list" {
		files, _ := os.Open(".")
		defer files.Close()

		fileInfo, _ := files.Readdir(-1)
		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	} else {
		fmt.Printf("Hello, %s!\n", *name)
	}
}
