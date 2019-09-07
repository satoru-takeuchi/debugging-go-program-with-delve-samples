package main

import (
	"fmt"
)

func main() {
	i := 1
	foo(i)
}

func foo(x int) {
	j := 2
	bar(x, j)
}

func bar(x int, y int) {
	k := 3
	fmt.Println(x, y, k)
}
