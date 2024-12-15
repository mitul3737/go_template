package main

import "fmt"

func main() {
	var city string = "Dhaka"
	fmt.Println("Welcome to", city, "City")
	fmt.Print("Hope to see you in", "\n", city)

	// only Print
	fmt.Print("Hope to see you in", city)

	//once we use println, we add a /n with the string actually
	fmt.Println("Lines are attched?")
	fmt.Println("On a different line?")

	//Printf -  format specifier
	var grades int = 42
	var name string = "Korim"
	fmt.Printf("My name is: %v \n", name)
	fmt.Printf("My marks in the current exam is: %d our of 45", grades)

}
