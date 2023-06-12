package main

import "fmt"

func main() {
	i := -10
	j := 25
	pI := &i
	pJ := &j
	fmt.Println("Value i, j", i, j)
	fmt.Println("Value i, j", *pI, *pJ)
	fmt.Println("Address i, j", pI, pJ)

	*pI = 12345
	*pI--
	fmt.Println("Value of i and address", i, pI)
	getPointer(pJ)
	fmt.Println("Change in value of J by reference", pJ, j)
	k := returnPointer(i)
	fmt.Println("square of i by passing value", k, *k)
}

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}
