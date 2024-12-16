package main

import "fmt"

// interface
// interface is a blueprint for a method set
// type <interface_name> interface{methods}
// can't introduce new variables rather work with methods
type shape interface {
	area() float64
	perimeter() float64
}

// struct
type square struct {
	side float64
}

// method
func (s square) area() float64 {
	return s.side * s.side
}

// method
func (s square) perimeter() float64 {
	return 4 * s.side
}

// struct
type rect struct {
	length, breadth float64
}

// method
func (r rect) area() float64 {
	return r.length * r.breadth
}

// method
func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

// method
func prindData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rect{length: 3, breadth: 4}
	c := square{side: 5}
	prindData(r)
	//{3 4}
	//12
	//14
	prindData(c)
	//{5}
	//25
	//20

}
