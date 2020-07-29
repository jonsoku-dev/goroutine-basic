package main

import (
	"fmt"
	"strings"
)

// return multiple value function
func lenAndUpper(name string) (length int, uppercase string) {
	// 함수가 끝나면 실행된다 "defer"
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	// naked return !
	// 사용자가 리턴값을 정해준대로 위에 변수가 존재하면 그대로 리턴한다.
	return
}

func main() {
	totalLength, _ := lenAndUpper("jongseok")
	fmt.Println(totalLength)
}
