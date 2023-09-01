package main

import (
	"fmt"
	"net/http"
	"time"
)

func sieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	var primeNumbers []int
	for i := 2; i <= n; i++ {
		if primes[i] {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	primes := sieveOfEratosthenes(10000000)
	elapsed := time.Since(start)
	fmt.Fprintf(w, "Number of primes: %d\nTime taken: %s\n", len(primes), elapsed)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
