package main

import (
	"fmt"
	"os"
)

func main() {
	//This is an example of how manipulate command as a data output with os package
	//to reproduce, its better generate a binary first, then in sequence, put the command
	//with arguments

	//comands to reproduce:
	//$ go build echo2.go
	//./echo2 a b c d 
	//output: a b c d
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
