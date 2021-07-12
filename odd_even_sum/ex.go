package main

import (
	"fmt"
)

const n int = 100

var even, odd int

func main() {

	even = 0
	odd = 0

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even += i
		} else {
			odd += i
		}
	}

	fmt.Println("even sum: ", even)
	fmt.Println("odd sum: ", odd)
}
