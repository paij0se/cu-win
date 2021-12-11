package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	switch {
	case len(os.Args) > 1:
		files, err := ioutil.ReadDir(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}

		for _, f := range files {
			fmt.Println(f.Name())
		}
	default:
		files, err := ioutil.ReadDir("./")
		if err != nil {
			fmt.Println(err)
		}

		for _, f := range files {
			fmt.Println(f.Name())
		}
	}

}
