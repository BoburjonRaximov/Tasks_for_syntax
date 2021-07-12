package main

import "fmt"

const n int = 10

func main() {
	var ar [n]int
	ar[0] = 0
	ar[1] = 1
	for i := 0; i < n; i++ {
		if i >= 0 && i < 2 {
			fmt.Println(ar[i])
		} else {
			ar[i] = ar[i-1] + ar[i-2]
			fmt.Println(ar[i])
		}
	}
}
