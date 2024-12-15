package main

import "fmt"

const PI float64 = 3.14 //constant value

func main() {
	//const <constant name> <constant type> = <value>
	const name string = "Harry Potter"
	const age int = 12
	const is_false = false

	fmt.Printf("%v %T\n", name, name) //value and type
	fmt.Printf("%v %T\n", is_false, is_false)
	fmt.Printf("%v %T\n", age, age)

	//area
	var radius float64 = 5.0
	var area float64
	area = PI * radius * radius
	fmt.Println("Area of Circle is: ", area)

}
