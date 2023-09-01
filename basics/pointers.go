package main

import "fmt"

func main() {
	var counter byte = 100
	P := &counter
	V := *P

	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

	V = 200
	fmt.Println()
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

	V = counter // en go todo se pasa como valor, no como referencia
	counter++
	fmt.Println()
	fmt.Printf("counter : %-16d address: %-16p\n", counter, &counter)
	fmt.Printf("P       : %-16p address: %-16p *P: %-16d\n", P, &P, *P)
	fmt.Printf("V       : %-16d address: %-16p\n", V, &V)

	// en sintesis
	// * antes de un tipo de dato es un puntero al tipo de dato, es decir guarda la direccion de memoria de ese tipo de dato
	// * antes de una variable es un operador de desreferencia, es decir, accede al valor de la direccion de memoria de la variable
	// & antes de una variable es un operador de direccion, es decir, accede a la direccion de memoria de la variable

	// actividad: crear 2 variables de tipos distintos, asignarles valores, crear 2 punteros y asignarles las direcciones de memoria de las variables
	// imprimir los valores de las variables, los punteros y los valores de los punteros
}
