package main

import (
	"fmt"
	"time"
)

func worker(id int, task int) int {
	fmt.Println("Worker", id, "processing task", task)
	time.Sleep(time.Second)
	return task * 2
}

func main() {
	start := time.Now()

	var results []int

	for j := 1; j <= 9; j++ {
		for i := 1; i <= 3; i++ {
			if j <= 9 {
				result := worker(i, j)
				results = append(results, result)
				j++
			}
		}
	}

	elapsed := time.Since(start)
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("Done")
	fmt.Printf("Time taken: %s\n", elapsed)
}
