package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) (result []string) {
	for i := 1; i <= n; i++ {
		isDivisibleBy3 := i%3 == 0
		isDivisibleBy5 := i%5 == 0

		if isDivisibleBy3 && isDivisibleBy5 {
			result = append(result, "FizzBuzz")
			continue
		} else if isDivisibleBy3 {
			result = append(result, "Fizz")
			continue
		} else if isDivisibleBy5 {
			result = append(result, "Buzz")
			continue
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return
}

func main() {
	fmt.Println("10")
	fmt.Println(fizzBuzz(10))

	fmt.Println("20")
	fmt.Println(fizzBuzz(20))

	fmt.Println("30")
	fmt.Println(fizzBuzz(30))

	fmt.Println("40")
	fmt.Println(fizzBuzz(40))

	fmt.Println("50")
	fmt.Println(fizzBuzz(50))
}
