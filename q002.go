// Even Fibonacci numbers
package main

import "fmt"

func main() {
	sum := 0
	prev := 0
	next := 1
	for {
		fibonacci := prev + next
		if fibonacci > 4000000 {
			break
		}
		if fibonacci%2 == 0 {
			sum += fibonacci
		}
		prev = next
		next = fibonacci
	}
	fmt.Println(sum)
	// sum = 4613732
}
