package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now())
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	// time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Difference: ", t1.Sub(t))
	format := t.Format("1 January 1970")
	fmt.Println("Formatted data", format)
	loc, _ := time.LoadLocation("America/New_York")
	americanTime := t.In(loc)
	fmt.Println("Time in USA", americanTime)

}
