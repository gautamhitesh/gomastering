package main

import "fmt"

type Digit int
type Power int

const (
	Zero Digit = iota
	One
	Two
	Three
	Four
)

const (
	p_0 Power = 1 << iota
	_
	p_2
	_
	p_4
	_
	p_6
)

func main() {
	fmt.Println("Digit type Zero", Zero)
	fmt.Println("Digit type One", One)
	fmt.Println("Digit type Two", Two)
	fmt.Println("Digit type three", Three)

	fmt.Println("Power type Zero", p_0)
	fmt.Println("Power type Two", p_2)
	fmt.Println("Power type Four", p_4)
	fmt.Println("Power type Six", p_6)

}
