// Sum square difference
package main

import "fmt"

func main() {
	num := 100
	diff := square_of_sum(num) - sum_of_squares(num) // <= 25164150
	fmt.Println(diff)
}

func sum_of_squares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}
	return sum
}

func square_of_sum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum * sum
}
