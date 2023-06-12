package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	// var myTime string
	myTime := os.Args[1]
	d, err := time.Parse("15:04:05", myTime) //15 04 05 //different constants for different date values
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Hour(), ":", d.Minute(), "::", d.Second())
	} else {
		fmt.Println(err)
	}
}
