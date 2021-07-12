package main

import (
	"fmt"
)

var a, b, c, d, e, f, g int

func main() {
	fmt.Println("to'rt xonali sonni kiriting")
	fmt.Scanln(&g)
	a = g / 1000 //1
	b = g / 100
	c = b % 10 // 2
	d = g / 10
	e = d % 10 //3
	f = g % 10 //4

	if g > 999 && g < 10000 {

		if a == f && c == e {
			fmt.Println("palindrome")

		} else {
			fmt.Println("no polindrome")
		}

	} else {
		fmt.Println("error")
	}

}
