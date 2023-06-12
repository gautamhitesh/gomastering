package main

import (
	"fmt"
)

//page 106
//declare 2D slice
//initialize using a loop

func main() {
	twoD := make([][]int, 3)
	counter := 1
	fmt.Println(twoD, len(twoD), cap(twoD))
	for i := 0; i < len(twoD); i++ {
		twoD[i] = make([]int, 3)
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = counter
			counter++
		}
	}
	//another way
	// initialize each row by appending each col
	twoD1 := make([][]int, 3)

	//this is needed because
	//default value for each slice is nil and hence each row is nil and should be initialized
	counter = 1
	for i := 0; i < len(twoD1); i++ {
		for j := 0; j < len(twoD1); j++ {
			twoD1[i] = append(twoD1[i], counter)
			counter++
		}
	}

	fmt.Println(twoD, twoD1)
	for rowNum, rowValue := range twoD1 {
		fmt.Println(rowNum, rowValue)
	}
}
