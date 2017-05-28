package main

import (
	"fmt"
	"log"
	"os"
)

// TODO:
// 1. find a way to strip out comments
// 2.

func main() {
	if len(os.Args) < 1 {
		// TODO: should we use a cli library here?
		log.Fatal("No file specified")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// TODO: is there a more idiomatic way to
	// read the complete file onto memory?
	prog := ""
	buf := make([]byte, 1)
	n := 1
	for n > 0 {
		n, err = file.Read(buf)
		if err != nil && n != 0 {
			log.Fatal(err)
		}
		// fmt.Println(buf)
		// fmt.Println(string(buf))
		prog += string(buf)
	}

	// TODO: why is there a '.' trailing the program
	fmt.Println(prog)

	// var data [30000]int
}
