package main

import (
	"fmt"
	"strings"
)

// return multiple value function
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("jongseok")
	fmt.Println(totalLength, upperName)
}
