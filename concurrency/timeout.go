package main

import (
	"fmt"
	"time"
)

func longRunningTask(ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- "Task completed"
}

func main() {
	ch := make(chan string)
	go longRunningTask(ch)

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(5 * time.Second):
		fmt.Println("Task timed out")
		// retry request api de un tercero
	}
}
