package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	fmt.Println(map1)
	map1[0] = "Hitesh"
	fmt.Println(map1)
	map1[1] = "Sumit"
	map1[2] = "Subham"
	delete(map1, 0)
	for key, value := range map1 {
		fmt.Println(key, ":", value)
	}

	present, ok := map1[0]
	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Doesn't exists", present)
	}
}
