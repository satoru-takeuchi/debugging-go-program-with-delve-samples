package main

import (
	"fmt"
)

func main() {
	n := 5
	sem1 := make(chan interface{})
	sem2 := make(chan interface{})
	fmt.Println("run five goroutines in parallel")
	for i := 0; i < n; i++ {
		go func(x int) {
			y := x * 2
			fmt.Println(y)
			sem1 <- 0
			<-sem2
		}(i)
	}
	for i := 0; i < n; i++ {
		<-sem1
	}
	for i := 0; i < n; i++ {
		sem2 <- 0
	}
	fmt.Println("done")
}
