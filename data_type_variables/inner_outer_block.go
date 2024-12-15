package main

import "fmt"

var name string = "Lisa" //global variable

func main() {
	// Outer block
	city := "Dhaka"
	{ // inner block
		country := "Bangladesh"
		fmt.Println(country)
		fmt.Println(city) //inner block can use content from outer block
	}

	fmt.Println(city)
	// can't use fmt.Println(country) as country variable is in the inner block

}
