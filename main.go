package main

import (
	asciiart "asciiart/test"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		file := asciiart.Spite("standard.txt")
		asciiart.PrintSymbole(file, os.Args[1])
	}

}
