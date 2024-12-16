package main

import "fmt"

// sturct
type Circle_1 struct {
	radius float64
	area   float64
}

// method
// here c is the pointer of the struct Circle_1 type
func (c *Circle_1) calcArea() {
	c.area = 3.14 * c.radius * c.radius //modifies the area of the object c
}

// struct  Student
type Student_0 struct {
	name   string
	grades []int
}

// method
func (s *Student_0) displayName() {
	fmt.Println(s.name)
}

// method
func (s *Student_0) calculatePercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func main() {
	c := Circle_1{radius: 5} //we have set radius only
	c.calcArea()             //once called calcArea
	fmt.Printf("%v", c)      //{5 78.5}

	//Student_1
	s := Student_0{name: "Joe", grades: []int{90, 75, 80}}
	s.displayName()                             //Joe
	fmt.Printf("%.2f", s.calculatePercentage()) //81.67
}
