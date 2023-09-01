package main

import (
	"fmt"
	"time"
)

func publisher(msg string, consumer func(string)) {
	consumer(msg)
}

func consumer(msg string) {
	fmt.Println("Received:", msg)
}

func main() {
	start := time.Now()

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Message %d", i+1)
		publisher(msg, consumer)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
