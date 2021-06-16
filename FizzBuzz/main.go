package main

import (
	"fmt"
	"strconv"
)

func main() {
	for num := 1; num <= 100; num++ {
		fmt.Println(fizzbuzz(num))
	}
}

func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"

	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(num)
	}
}
