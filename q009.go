// Special Pythagorean triplet
package main

import "fmt"

func main() {
	a := 0
	b := 0
	c := 0
	for i := 1000; i > 0; i-- {
		c = i
		for j := 1; j < 1000-c; j++ {
			b = j
			a = 1000 - c - b
			if a > 0 {
				if a*a+b*b == c*c {
					i = 0
					break
				}
			}
		}
	}
	fmt.Println(a * b * c) // a = 200 b = 375 c = 425
}
