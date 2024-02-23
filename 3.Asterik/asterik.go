package main

import "fmt"

func playWithAsterik(number int) {
	for i := 0; i < number; i++ {
		for k := 0; k <= i; k++ {
			fmt.Print("*")

		}
		fmt.Println()
	}
}

func main() {
	playWithAsterik(5)
	// *
	// **
	// ***
	// ****
	// *****

	playWithAsterik(7)
	// *
	// **
	// ***
	// ****
	// *****
	// ******
	// *******
}
