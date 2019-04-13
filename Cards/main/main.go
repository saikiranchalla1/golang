package main

import (
	"fmt"
)

func main() {
	var numbers []int

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for j := range numbers {
		if numbers[j]%2 == 0 {
			fmt.Println(numbers[j], " is even")
		} else {
			fmt.Println(numbers[j], " is odd")
		}
	}
}
