package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// basic function call with no return values
	welcome()

	// functions can return values
	// name := "ppp"
	// check := welcomeByName(name)
	// if check {
	// 	fmt.Println("Welcome,", name)
	// } else {
	// 	println(name, "Your name is not on the guest list!!!")
	// }

	// Pass a name from CLI and then check if that name is welcome or not
	guest := os.Args[1]
	check := welcomeByName(os.Args[1])
	if check {
		fmt.Println("Welcome,", guest)
	} else {
		println(guest, "Your name is not on the guest list!!!")
	}

	// ... means that give this array the size as per count of values
	scores := []int{10, 20, 30, 40}
	length, total := getSizeAndCount(scores)
	fmt.Println("Length: ", length)
	fmt.Println("Total: ", total)

}

// func getSizeAndCount(s []int) (int, int) {
// 	size := len(s)
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	return size, sum
// }

func getSizeAndCount(s []int) (size int, sum int) {
	size = len(s)
	sum = 0
	for _, v := range s {
		sum += v
	}
	return
}

func welcome() {
	fmt.Println("Welcome to Uranus!")
}

func welcomeByName(nm string) bool {

	result := strings.Contains(nm, "a")
	return result

	// return true or return false
	// if any of the result1, result2, result3, result4, result5 is true then return true
	// return true if any vowels are in the name, else return false
	result1 := strings.Contains(nm, "a")
	result2 := strings.Contains(nm, "e")
	result3 := strings.Contains(nm, "i")
	result4 := strings.Contains(nm, "o")
	result5 := strings.Contains(nm, "u")

	if result1 || result2 || result3 || result4 || result5 {
		return true
	} else {
		return false
	}

}
