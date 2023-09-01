package main

import (
	"fmt"
	"time"
)

func largeProcessing() {
	time.Sleep(time.Second * 3)
	fmt.Println("Terminó el procesamiento")
}

func main() {
	go largeProcessing() // Se ejecuta en una goroutine, es decir en otro hilo     X<---------x
	fmt.Println("Procesando...")
	fmt.Println("Done")
	time.Sleep(time.Second * 4)
}
