package main

import "fmt"

func string_demo() {

	s := "hello, world"

	hello := s[:5]

	world := s[7:]

	fmt.Println(hello, world)

	for i, v := range []byte("世界abcd") {
		fmt.Printf("序号 %d: %b;", i, v)
	}
}
