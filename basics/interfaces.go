package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle of width %f and height %f", r.Width, r.Height)
}

type Shape interface {
	Area() float64      // no params and returns a float64
	Perimeter() float64 // no params and returns a float64
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func Describe(i interface{}) { // any type
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

// composition
type Reader interface {
	Read(b []byte) (int, error)
}

type Writer interface {
	Write(b []byte) (int, error)
}

type ReaderWriter interface {
	Reader
	Writer
}

func main() {
	rectangulo := Rectangle{10, 5}
	circulo := Circle{10}
	area1 := CalculateArea(rectangulo)
	area2 := CalculateArea(circulo)
	fmt.Println("Area:", area1)
	fmt.Println("Area:", area2)
	Describe(nil)
}

// 1. Crear una interfaz llamada Talker: Esto significa que debes definir un nuevo tipo de interfaz llamado Talker que describa objetos que pueden "hablar".
// 2. Con un único método Talk() string: Esto significa que cualquier tipo que desee implementar esta interfaz Talker debe tener un método llamado Talk que devuelva una cadena (string).
// 3. Luego implementa esta interfaz para dos estructuras diferentes: Human y Robot: Necesitas definir dos estructuras diferentes, una llamada Human y otra llamada Robot. Ambas estructuras deben tener un método Talk() que devuelva un string, cumpliendo así con la interfaz Talker.
// 4. Finalmente, escribe una función que acepte un Talker y lo haga hablar: Crea una función que tome un argumento de tipo Talker (que podría ser un Human o un Robot, ya que ambos implementan Talker). Esta función debe llamar al método Talk() en su argumento e imprimir la cadena devuelta.
