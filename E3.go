package main

import "fmt"

func main() {
	anArray := make([]int, 7)
	counter := 1
	for key := range anArray {
		anArray[key] = counter
		counter++
	}
	fmt.Println(anArray)
	aMap := make(map[int]int)
	for key, value := range anArray {
		aMap[key] = value
	}
	fmt.Println(aMap)
}
