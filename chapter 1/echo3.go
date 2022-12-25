package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//This is an example of how manipulate command as a data output with os package
	//to reproduce, its better generate a binary first, then in sequence, put the command
	//with arguments

	//comands to reproduce:
	//$ go build echo3.go
	//./echo3 a b c d
	//output: a b c d
	
	fmt.Println(strings.Join(os.Args[1:], " "))
}
