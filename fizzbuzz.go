package main

import (
		"fmt"
		"strconv"
)

func main () {
	fmt.Println("startmain")
}

func Fizzbuzz (number int) string {

	if (number % 15) == 0 {
		return "fizzbuzz"
	}
	if (number % 3) == 0 {
		return "fizz"
	}
	if (number % 5) == 0 {
		return "buzz"
	}

	return strconv.Itoa(number)
}