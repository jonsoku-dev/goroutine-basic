package main

import (
	"fmt"
)

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("jongseok", "kazuko", "hiromi", "mother")
}
