package main

import "fmt"

//func <func_name>(variable_1 variable_type_1, variable_2, variable_type_2)return type{}

func addNumbers(a int, b int) int {
	sum := a + b
	return sum //returns an integer
}

func printGreetings(str_1 string) {
	fmt.Println("Hey there", str_1)
}

func operations(a int, b int) (int, int) {
	sum_0 := a + b
	diff_0 := a - b
	return sum_0, diff_0
}

//variadic function
//func <func_name> (param_1 type, param_2 type, param_3 type, param_4 ...type)<return_type>

func variadic_sum(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func varia_string(student string, subjects ...string) {
	fmt.Println("hey", student, ", here are your subjects -")
	for _, sub := range subjects {
		fmt.Printf("%s\n", sub)
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

// defer
func f1(s string) {
	fmt.Println(s)
}
func f2(k string) {
	fmt.Println(k)
}
func f3(l int) {
	fmt.Println(l)
}
func f4(lp string) {
	fmt.Println("Done with  everything!", lp)
}

func main() {
	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)
	printGreetings("Bro!")
	fmt.Println(operations(9, 10))

	//this function takes whatever amount of input we want to give
	fmt.Println(variadic_sum())
	fmt.Println(variadic_sum(10))
	fmt.Println(variadic_sum(10, 20))
	fmt.Println(variadic_sum(100, 90, 30))

	varia_string("Sakib", "Bangla", "English")

	//factorial
	fmt.Println(factorial(5))

	//anonymous function
	x := func(l int, b int) int {
		return l * b
	}
	fmt.Printf("%T\n", x)
	fmt.Println(x(20, 30))

	//practicing defer one
	f1("Hey bro")
	defer f2("Last e hobe") //this will word  once other functions are done
	f3(1234)
	f4("KORIM")

}
