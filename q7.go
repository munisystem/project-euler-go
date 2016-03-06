// 10001st prime
package main

import "fmt"

func main() {
	fmt.Println(get_prime(10001)) // <= 104743
}

func get_prime(num int) int {
	primes := []int{2}
	for i := primes[len(primes)-1]; num != len(primes); i++ {
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
	return primes[num-1]
}
