package main

import "fmt"

// Define the Person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) FullName() {
	fmt.Println(p.FirstName, p.LastName)
}

func main() {
	// p := Person{"John", "Doe", 30}

	// fmt.Println("First Name:", p.FirstName)
	// fmt.Println("Last Name:", p.LastName)
	// fmt.Println("Age:", p.Age)

	// p.FullName()

	// p1 := Person{
	// 	FirstName: "Alexis",
	// 	LastName:  "Sanchez",
	// 	Age:       30,
	// }
	// p1.FullName()

	// fmt.Println(p1)
	// fmt.Printf("%#v\n", p1)
	// fmt.Printf("Tipo de p1: %T\n", p1)
	// fmt.Println(p1.FirstName)

	// composicion de structs  (composition)

	type Employee struct {
		Person
		JobTitle string
	}

	e1 := Employee{
		Person: Person{
			FirstName: "Alexis",
			LastName:  "Sanchez",
			Age:       30,
		},
		JobTitle: "Developer",
	}

	fmt.Println(e1)
	fmt.Printf("%#v\n", e1)
	fmt.Printf("Tipo de e1: %T\n", e1)

	fmt.Println(e1.Person.FirstName)

	type PersonalityDisorder struct {
		Person1 Person
		Person2 Person
		People  []Person
	}
}

// https://github.com/inancgumus/learngo/blob/master/24-structs/exercises/01-warmup/main.go
