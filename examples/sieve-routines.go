package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func handlerWithGoroutines(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var primes []int
	var wg sync.WaitGroup

	for i := 2; i <= 100000; i += 10000 {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i <= end; i++ {
				isPrime := true
				for j := 2; j*j <= i; j++ {
					if i%j == 0 {
						isPrime = false
						break
					}
				}
				if isPrime {
					primes = append(primes, i)
				}
			}
		}(i, i+9999)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Fprintf(w, "Number of primes: %d\nTime taken: %s\n", len(primes), elapsed)
}

func main() {
	http.HandleFunc("/", handlerWithGoroutines)
	http.ListenAndServe(":8091", nil)
}
