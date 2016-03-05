// Largest prime factor
package main

import "fmt"

func main() {
	target := 600851475143
	for i := 2; i < target/2; i++ {
		for target%i == 0 {
			target = target / i
		}
	}
	fmt.Println(target)
}
