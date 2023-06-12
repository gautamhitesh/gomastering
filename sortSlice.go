package main

import (
	"fmt"
	"sort"
)

// page 109
// sorting of a compilcated data type
type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	structSlice := make([]aStructure, 0)
	structSlice = append(structSlice, aStructure{"Hitesh", 178, 112})
	structSlice = append(structSlice, aStructure{"Sumit", 179, 75})
	structSlice = append(structSlice, aStructure{"Subham", 173, 62})
	fmt.Println("Unsorted ", structSlice)
	//sort needs a less function which should return true for less
	sort.Slice(structSlice, func(i, j int) bool { return structSlice[i].height < structSlice[j].height })
	fmt.Println("Sorted by height", structSlice)
	sort.Slice(structSlice, func(i, j int) bool { return structSlice[i].weight < structSlice[j].weight })
	fmt.Println("Sorted by weight", structSlice)

}
