package main

// Para paralelizar el trabajo entre multiples trabajadores y limitar el maximo de trabajos encolados
// Por ejemplo ppara procesar multiples ordenes de compra
import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for t := range tasks {
		fmt.Println("Worker", id, "processing task", t)
		time.Sleep(time.Second)
		results <- t * 2
	}
}

func main() {
	start := time.Now()
	tasks := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	for j := 1; j <= 9; j++ {
		tasks <- j
	}
	close(tasks)

	for k := 1; k <= 9; k++ {
		fmt.Println(<-results)
	}

	fmt.Println("Done")
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
