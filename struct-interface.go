package main

import (
	"fmt"
)

type position struct {
	x int
	y int
}

type positionInterface interface {
	X() int
	Y() int
}

func (pos position) X() int {
	return pos.x
}

func (pos position) Y() int {
	return pos.y
}

func main() {
	pos := position{
		x: 1,
		y: 2,
	}
	var posIntf positionInterface = pos
	fmt.Println(pos, posIntf)
}
