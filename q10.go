package main

import "fmt"

func main() {
	num := 2000000
	primes := below_primes(num)
	sum := 0
	for _, v := range primes {
		sum += v
	}
	fmt.Println(sum) //<= 142913828922
}

func below_primes(num int) []int {
	primes := []int{2}
	for i := primes[len(primes)-1] + 1; num > i; i++ {
		flag := 0
		for _, v := range primes {
			if i%v == 0 {
				flag++
			}
		}
		if flag == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}
