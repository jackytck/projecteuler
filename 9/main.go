package main

import "fmt"

func main() {
	n := 1000
	for a := 1; a < n; a++ {
		for b := 1; b < n; b++ {
			for c := 1; c < n; c++ {
				if a+b+c == n && a*a+b*b == c*c {
					fmt.Println(a*b*c, a, b, c)
					return
				}
			}
		}
	}
}

// Find abc such that a^2+b^2=c^2 and a+b+c=1000.
