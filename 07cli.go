package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println(os.Args)
	fmt.Println(os.Args[1:])
	// fmt.Println(reflect.TypeOf(os.Args[1]))
	sum := 0
	for _, v := range os.Args[1:] {
		// fmt.Println(index, ".", v)
		// sum = sum + v
		fmt.Println(v)
		// this is ignoring error or invalid input
		value, _ := strconv.Atoi(v)

		// // if you want to incorporate input validation
		// value, error := strconv.Atoi(v)
		// if error != nil {
		// 	fmt.Println(error)
		// 	os.Exit(1)
		// }

		sum = sum + value
	}
	fmt.Println("Sum: ", sum)
}
