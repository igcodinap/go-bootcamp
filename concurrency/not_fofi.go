package main

import (
	"fmt"
	"time"
)

func square(x int) int {
	return x * x
}

func main() {
	start := time.Now()

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var results []int

	for _, n := range numbers {
		results = append(results, square(n))
	}

	elapsed := time.Since(start)
	fmt.Println("Results:", results)
	fmt.Printf("Time taken: %s\n", elapsed)
}
