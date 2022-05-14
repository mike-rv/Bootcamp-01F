package main

import "fmt"

func isPrime(num int) bool {
	var div int
	// Count the divisors other than 1 and
	// num.
	for i := 2; i < num; i++ {
		if num%i == 0 {
			div++
		}
	}

	return div == 0
}

func reportPrime(num int) {
	fmt.Println("Is", num, "prime?")
	var result bool = isPrime(num)
	fmt.Println(result)

}

func main() {
	fmt.Println("Prime numbers")
	reportPrime(619)
	reportPrime(319)
	reportPrime(519)
	reportPrime(419)

}
