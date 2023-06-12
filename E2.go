package main

import (
	"fmt"
	"time"
)

const (
	day0 = time.Weekday(iota)
	day1
	day2
	day3
	day4
	day5
	day6
)

func main() {
	fmt.Println(day0, day1, day2, day3, day4, day5, day6)
}
