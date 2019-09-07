package main

import (
	"fmt"
)

func main() {
	res := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; i++ {
			res += i * j
		}
	}
	fmt.Println(res)
}
