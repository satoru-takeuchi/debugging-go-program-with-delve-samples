package main

import (
	"fmt"
)

func main() {
	const n = 10
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = i * 2
	}
	fmt.Println(a)
}
