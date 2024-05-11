package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().Unix())

	switch num := rand.Intn(20); {
	case num > 15:
		fmt.Printf("成功啦! 当前随机数为 %d\n", num)

	case num < 15:
		fmt.Printf("失败了，呜呜呜 当前随机数为 %d\n", num)

	default:
		fmt.Printf("天选之人！ 当前随机数为 %d\n", num)
	}
}
