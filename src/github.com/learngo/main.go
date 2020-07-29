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

func superAdd(numbers ...int) int {
	// index가 처음
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
