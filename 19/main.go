package main

import (
	"fmt"
	"time"
)

func countWeekday(fromYear, toYear int, weekday time.Weekday) int {
	var cnt int
	for y := fromYear; y <= toYear; y++ {
		for m := 1; m <= 12; m++ {
			d := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
			if d.Weekday() == weekday {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countWeekday(1901, 2000, 0))
}

// Count the number of Sunday on the first date of the month.
