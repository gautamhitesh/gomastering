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

	myTime := os.Args[1]
	d, err := time.Parse("Jan 2, 2006, Mon, 15:04:05", myTime)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Hour(), ":", d.Minute(), "::", d.Second())
	} else {
		fmt.Println(err)
	}

}

//E5 and E6 similar
//use regex E5
