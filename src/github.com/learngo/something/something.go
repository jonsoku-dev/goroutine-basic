package something

import "fmt"

func sayBye() {
	fmt.Println("bye")
}

// 대문자인 경우에는 export 의 의미를 갖는다.
func SayHello() {
	fmt.Println("Hello")
}
