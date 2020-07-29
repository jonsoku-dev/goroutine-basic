package main

import (
	"fmt"
)

// 둘다 인트면 나중에 적어도 인식한다.
func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}
