package main

// Para desacoplar productores y consumidores
// Por ende, sirve para la implementacion de arquitecturas en base a eventos
import (
	"fmt"
	"time"
)

type PubSub struct {
	subscribers []chan int
}

func (ps *PubSub) Publish(value int) {
	for _, ch := range ps.subscribers {
		ch <- value
	}
}

func (ps *PubSub) Subscribe() chan int {
	ch := make(chan int)
	ps.subscribers = append(ps.subscribers, ch)
	return ch
}

func main() {
	start := time.Now()

	ps := &PubSub{}

	sub1 := ps.Subscribe()
	go func() {
		for i := range sub1 {
			fmt.Println("Subscriber 1 received:", i)
		}
	}()

	sub2 := ps.Subscribe()
	go func() {
		for i := range sub2 {
			fmt.Println("Subscriber 2 received:", i)
		}
	}()

	ps.Publish(1)
	ps.Publish(2)
	fmt.Println("Done")
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
