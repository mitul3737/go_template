package main

import "fmt"

// struct to keep same things together
// type <struct name> struct{ list of fields}
type Student struct {
	name   string
	rollNo int
	//marks  []int
	//grades map[string]int
}

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

// using passing by reference concept as passing by value does not change struct
func calcArea(c *Circle) {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	(*c).area = area
}

type s1 struct {
	x int
}

func main() {
	var s Student
	fmt.Printf("%+v", s) //output {name: rollNo:0 marks:[] grades:map[]} for s

	//struct initialization
	//variable_name:=new(<struct_name>)
	st := new(Student)
	//adds new values for student with &
	fmt.Printf("%v", st) //{name: rollNo:0 marks:[] grades:map[]}&{ 0 [] map[]}
	//here last { 0 [] map[]} is for st

	//initialization
	/*variableName:=<struct_name>
	field)name: value,
	field_name; value.*/

	st_1 := Student{"Shakib", 75}
	fmt.Printf("%v", st_1) //{name: rollNo:0}&{ 0}{Shakib 75}
	//here last { 0}{Shakib 75} is for the st_1

	var c Circle
	c.x = 10
	c.y = 5
	c.radius = 5
	fmt.Printf("%v \n", c) //gives x,y,radius and area {10 5 5 0}

	//using passing by reference
	c_1 := Circle{x: 5, y: 5, radius: 5, area: 0}
	fmt.Printf("%v\n", c_1) //remains same because works passing by value
	calcArea(&c_1)
	fmt.Printf("%v\n", c_1) //changes because works as passing by reference

	//comparing struct
	m_0 := s1{x: 5}
	m_1 := s1{x: 19}
	//as both m_0 and m_1 are of same type s1 , we can compare
	fmt.Println(m_0 == m_1) //5 is not equal 19

}
