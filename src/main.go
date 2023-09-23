package main

import (
	"fmt"
	"os"
)
func printOsArgs () {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	println("main running")
	printOsArgs()
}   
