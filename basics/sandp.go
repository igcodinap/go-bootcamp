package main

import "fmt"

type Person struct {
	Age int
}

func (p Person) SetAgeValueReceiver(age int) Person {
	fmt.Printf("SetAgeValueReceiver: %p\n", &p)
	p.Age = age
	fmt.Println("SetAgeValueReceiver:", p.Age)
	return p
}

func (p *Person) SetAgePointerReceiver(age int) {
	fmt.Printf("SetAgePointerReceiver: %p\n", p)
	p.Age = age
	fmt.Println("SetAgePointerReceiver:", p.Age)
}

func main() {
	person := Person{Age: 30}

	p2 := person.SetAgeValueReceiver(35)
	fmt.Println("Outside SetAgeValueReceiver:", person.Age)

	person.SetAgePointerReceiver(35)
	fmt.Println("Outside SetAgePointerReceiver:", person.Age)
}
