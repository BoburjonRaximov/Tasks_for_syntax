package main

import (
	"fmt"
)

const n int = 30

func main() {

	for i := 0; i < n; i++ {
		if i%15 == 0 && i > 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 && i > 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 && i > 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
