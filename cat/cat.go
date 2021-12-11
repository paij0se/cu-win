package main

import "os"

func main() {
	switch {
	case len(os.Args) > 1:
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		b := make([]byte, 100)
		for {
			n, err := f.Read(b)
			if err != nil {
				break
			}
			os.Stdout.Write(b[:n])
		}
	default:
		os.Stdout.WriteString("Usage: cat <filename>\n")

	}
}
