package main

import (
	asciiart "asciiart/func"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	} else {
		file := asciiart.Splite("standard.txt")
		asciiart.PrintSymbole(file, os.Args[1])
	}

}
