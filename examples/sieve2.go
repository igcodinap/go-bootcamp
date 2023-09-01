package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sieveSegment(start, end int, primes []bool) {
	for p := 2; p*p <= end; p++ {
		for i := max(p*p, (start+p-1)/p*p); i <= end; i += p {
			primes[i-start] = false
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	n := 10000000
	numSegments := 4
	segmentSize := n / numSegments

	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	var wg sync.WaitGroup

	for i := 0; i < numSegments; i++ {
		wg.Add(1)
		go func(i int) {
			start := i * segmentSize
			end := (i + 1) * segmentSize
			if i == numSegments-1 {
				end = n
			}
			sieveSegment(start, end, primes[start:end+1])
			wg.Done()
		}(i)
	}

	wg.Wait()

	var primeNumbers []int
	for i := 2; i <= n; i++ {
		if primes[i] {
			primeNumbers = append(primeNumbers, i)
		}
	}

	elapsed := time.Since(start)
	fmt.Fprintf(w, "Number of primes: %d\nTime taken: %s\n", len(primeNumbers), elapsed)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
