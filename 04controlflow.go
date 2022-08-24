package main

import "fmt"

func main() {
	score := 88
	if score >= 88 {
		fmt.Printf("%v is greater than or equal to 88.\n", score)
	}

	if score >= 90 {
		fmt.Printf("%v is greater than or equal to 90.\n", score)
	} else {
		fmt.Printf("%v is not greater than 90.\n", score)
	}
}
