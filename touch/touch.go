package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("touch <file>")
		os.Exit(1)
	}
	myfile, e := os.Create(os.Args[1])
	if e != nil {
		fmt.Println(e)
	}
	log.Println("created file:", myfile.Name())
	myfile.Close()
}
