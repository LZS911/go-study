package main

import "fmt"

func array_demo() {
	var a [3]int // 定义长度为 3 的 int 型数组, 元素全部为 0

	var b = [...]int{1, 2, 3} // 定义长度为 3 的 int 型数组, 元素为 1, 2, 3

	var c = [...]int{2: 3, 1: 2} // 定义长度为 3 的 int 型数组, 元素为 0, 2, 3

	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为 6 的 int 型数组, 元素为 1, 2, 0, 0, 5, 6

	fmt.Println(a, b, c, d)

	var e = &d

	for i, v := range a {
		fmt.Println(i, v)
	}

	for i, v := range e {
		fmt.Println(i, v)
	}

	var time [5][0]int

	for range time {
		fmt.Println("hello world")
	}
}
