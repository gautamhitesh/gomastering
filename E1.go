package main

import "fmt"

const (
	pow4_0 = 1 << (2 * iota)
	pow4_1
	pow4_2
	pow4_3
)

const (
	pow2_0 = 1 << iota
	pow2_1
	pow2_2
	pow2_3
)

func main() {
	fmt.Println(pow4_0, pow4_1, pow4_2, pow4_3)
	fmt.Println(pow2_0, pow2_1, pow2_2, pow2_3)
}
