package main

import "fmt"

func main() {
	numbers := []int{1, 3, 4, 5, 6, 7}
	users := map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	//For loops
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	for _, i := range numbers {
		fmt.Println(i)
	}

	for k, v := range users {
		fmt.Println(k, v)
	}
}
