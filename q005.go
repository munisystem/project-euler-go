// Smallest multiple
package main

import "fmt"

func main() {
	mult := 1
	end := 20
	for i := 2; i <= end; i++ {
		if is_prime(i) == true {
			mult = mult * largest_factorial(i, end)
		}
	}
	fmt.Println(mult)
}

func is_prime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func largest_factorial(num int, end int) int {
	fuct := num
	for {
		if fuct*num > end {
			return fuct
		}
		fuct = fuct * num
	}
}
