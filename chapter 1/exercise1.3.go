package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	compareEcho1(os.Args)
	compareEcho2(os.Args)
	compareEcho3(os.Args)
}

	//This is an example of how manipulate command as a data output with os package
	//to reproduce, its better generate a binary first, then in sequence, put the command
	//with arguments

	//comands to reproduce:
	//$ go build exercise1.3.go
	//exercise1.3 a b c d
	//output: 
	//a b c d 4.522µs
	//a b c d 1.001µs
	//a b c d 663ns

func compareEcho1(args []string) {
	start := time.Now()
	var s, sep string

	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	end := time.Now()

	fmt.Println(s, end.Sub(start))
}

func compareEcho2(args []string) {
	start := time.Now()
	s, sep := "", ""

	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	end := time.Now()

	fmt.Println(s, end.Sub(start))
}

func compareEcho3(args []string) {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "), time.Now().Sub(start))
}
