package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// dup1()
	// dup2()
	// dup3()
	dup4()
}

func dup1() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("出现重复：%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		count_lines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			count_lines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("出现重复：%d\t%s\n", n, line)
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("出现重复：%d\t%s\n", n, line)
		}
	}
}

func count_lines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}

type Info struct {
	count    int
	filename string
}

func dup4() {
	counts := make(map[string]*Info)
	files := os.Args[1:]

	if len(files) == 0 {
		count_lines_2(os.Stdin, counts, "stdio")
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			count_lines_2(f, counts, filename)

			f.Close()
		}
	}
	for line, info := range counts {
		if info.count > 1 {
			fmt.Printf("出现重复：%d\t%s\t%s\n", info.count, info.filename, line)
		}
	}
}
func count_lines_2(f *os.File, counts map[string]*Info, filename string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = &Info{count: 1, filename: ""}
		} else {
			counts[input.Text()].count = counts[input.Text()].count + 1
			counts[input.Text()].filename = filename
		}

	}
}
