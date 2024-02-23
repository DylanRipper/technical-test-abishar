package main

import "fmt"

func recursiveFactorial(num int) int {
	if num > 0 {
		return num * recursiveFactorial(num-1)
	}

	return 1
}

func main() {
	fmt.Println(recursiveFactorial(5)) // 120
	fmt.Println(recursiveFactorial(7)) // 5040
	fmt.Println(recursiveFactorial(9)) // 362880
	fmt.Println(recursiveFactorial(3)) // 6
}
