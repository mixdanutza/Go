package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct{}

type triangle struct{}

func main() {
	t := triangle{}
	s := square{}

	printArea(t)
	printArea(s)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return 4.4
}

func (t triangle) getArea() float64 {
	return 6
}
