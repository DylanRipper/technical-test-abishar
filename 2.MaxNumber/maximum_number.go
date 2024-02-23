package main

import "fmt"

func CheckMaximumNumber(arrNum []int) int {
	var max int
	for _, num := range arrNum {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	fmt.Println(CheckMaximumNumber([]int{3, 5, 1, 9, 2}))              // 9
	fmt.Println(CheckMaximumNumber([]int{9, 10, 3, 7, 4, 2, 6}))       // 10
	fmt.Println(CheckMaximumNumber([]int{97, 60, 38, 72, 49, 20, 56})) // 97
}
