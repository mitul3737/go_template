package main

import (
	"fmt"
	"reflect" //for type
	"strconv" //sting to Integer
)

func main() {
	var name string
	var age int
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s %d", &name, &age) //input can be : Korim 29
	fmt.Println("Your name is  ", name, " and ", age, " for sure!")
	//output can be: Your name is  Korim and 29 for sure!

	//count ; number of arguments that the function writes to
	// err any error thrown during the execution of the function
	/*var a string
	var b int
	fmt.Print("Enter a string and a number")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count :", count)
	fmt.Println("error :", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)*/

	//Type
	fmt.Printf("Your name is %v of type %T and age %v of type %T for sure!", name, name, age, age)
	fmt.Printf("\n")

	fmt.Printf("Type %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type %v \n", reflect.TypeOf("Shahriyar"))

	//convert type
	var i int = 90
	var f float64 = float64(i) //convert to float
	fmt.Printf("%.3f\n", f)    //3 values to  print after .  decimal point

	// Integer to string
	var ik int = 42
	var sp string = strconv.Itoa(ik) //convert to string
	fmt.Printf("%q\n", sp)
	//string to integer
	var ck string = "200"
	mk, err := strconv.Atoi(ck)
	fmt.Printf("%v, %T \n", mk, mk)   //value and error
	fmt.Printf("%v, %T \n", err, err) //no error was there

	var lk string = "299abcd"
	wk, err := strconv.Atoi(lk)
	fmt.Printf("%v, %T\n", wk, wk)
	fmt.Printf("%v, %T\n", err, err) //showing error as we can't convert 299abcd into any string. So, err variable stores the error

}
