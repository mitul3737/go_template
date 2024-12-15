package main

import "fmt"

func main() {

	//same type variable
	var s, t = "foo", "bar"
	fmt.Println(s)
	fmt.Println(t)

	//different type variable
	var (
		st string = "foo"
		i  int    = 5
	)
	fmt.Println(st)
	fmt.Println(i)

	//short variable assignment
	sk := "Hello World" //assigned
	sk = "Hacked it!"
	fmt.Println(sk)
}
