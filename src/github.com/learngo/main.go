package main

// return multiple value function
// func lenAndUpper(name string) (length int, uppercase string) {
// 	// 함수가 끝나면 실행된다 "defer"
// 	defer fmt.Println("I'm done!")
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	// naked return !
// 	// 사용자가 리턴값을 정해준대로 위에 변수가 존재하면 그대로 리턴한다.
// 	return
// }

// func superAdd(numbers ...int) int {
// 	// index가 처음
// 	// for index, number := range numbers {
// 	// 	fmt.Println(index, number)
// 	// }
// 	// for i := 0; i < len(numbers); i++ {
// 	// 	fmt.Println(numbers[i])
// 	// }
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func canIDrink(age int) bool {
// 	if koreanAge := age + 2; koreanAge < 18 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

// func canIDrinkWithSwitch(age int) bool {
// 	switch koreanAge := age + 2; koreanAge {
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }

/**
	// memory copy
	a := 2
	b := &a -> memory값을 저장
	fmt.Print(&a, &b)
	fmt.Print(*b) -> memory값 해석
**/

/**
	// [5] 라는 배열의 값을 정해주었다 new Array(5) 와 같음
	names := [5]string{"jongseok", "kazuko", "takako"}
	names[3] = "111"
	names[4] = "111"
	names[5] = "111" // <-- 5를 초과하였으니 error!
	fmt.Println(names)

	// [] 를 비워두면 뒤에 값이 오는만큼 채워짐
**/

/**
	names := []string{"jongseok", "kazuko", "takako"}
	// names[3] = "lala" // <-- 3이 비어있으므로 에러가난다.
	names = append(names, "fuck")
	fmt.Println(names)
**/

/**
	// struct
	// map[key]value{key:value, key:value}
	jongseok := map[string]string{"name": "jongseok", "age": "19"}
	fmt.Println(jongseok)
**/

/**
	jongseok := map[string]string{"name": "jongseok", "age": "19"}
	for key, value := range jongseok {
		fmt.Println(key, value)
	}
**/

/**
	// struct
	type person struct {
		name    string
		age     int
		favFood []string
	}

	func main() {
		favFood := []string{"kimchi", "sundubu"}
		jongseok := person{"jongseok", 18, favFood}
		fmt.Println(jongseok)
	}
**/
/**
	// struct (better than top)
	type person struct {
		name    string
		age     int
		favFood []string
	}

	func main() {
		favFood := []string{"kimchi", "sundubu"}
		jongseok := person{name: "jongseok", age: 18, favFood: favFood}
		fmt.Println(jongseok)
	}
**/
func main() {

}
