package main

import "fmt"

func CheckPalindrome(word string) bool {
	var temp string
	for i := len(word) - 1; i >= 0; i-- {
		temp += string(word[i])
	}

	return temp == word
}

func main() {
	fmt.Println(CheckPalindrome("tamat")) // true
	fmt.Println(CheckPalindrome("katak")) // true
	fmt.Println(CheckPalindrome("motor")) // false
	fmt.Println(CheckPalindrome("radar")) // true
	fmt.Println(CheckPalindrome("hello")) // false
}
