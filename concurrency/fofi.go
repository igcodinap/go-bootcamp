package main

// Para paralelizar el trabajo entre multiples consumidores y posteriormente procesar los resultados de cada uno juntos
// Por ejemplo para procesar un archivo de 1GB en 10 partes de 100MB cada una
import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func consumer(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func main() {
	start := time.Now()

	numbers := make(chan int)
	squares := make(chan int)

	go producer(numbers)
	go consumer(numbers, squares)
	fmt.Println("Results:")

	for s := range squares {
		fmt.Println(s)
	}

	fmt.Println("Done")
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
