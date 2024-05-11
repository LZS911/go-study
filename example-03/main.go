package main

import (
	"fmt"
	"strings"

	"rsc.io/quote"

	"os"
)

func main() {
	fmt.Println(os.Args[0])

	s, sep := "", ""

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)

		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[1:])

	fmt.Println(quote.Go())
}
