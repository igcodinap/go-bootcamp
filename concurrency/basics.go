package main

import (
	"fmt"
	"time"
)

// Esta función se ejecutará en su propia goroutine
func printNumbers(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i                            // envía el valor de i al canal ch
		time.Sleep(time.Millisecond * 100) // espera 100 milisegundos
	}
	close(ch) // cierra el canal
}

func main() {
	// Crea un nuevo canal de enteros
	myChannel := make(chan int)

	// Inicia la función printNumbers en una nueva goroutine
	go printNumbers(myChannel)

	// Recibe valores del canal
	for num := range myChannel {
		fmt.Println("Recibido:", num)
	}
}
