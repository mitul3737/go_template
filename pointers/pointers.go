package main

import "fmt"

func modify_by_value(sk string) {
	sk = "world"                    //saved in 0xc0000140e0                    //saved in an address like 0xc000114050
	fmt.Printf("%T %v\n", sk, sk)   // the value showws world only
	fmt.Printf("%T %v\n", &sk, &sk) //address type and address value
}

func modify_by_reference(sm *string) {
	*sm = "world1"                  //saved in 0xc000128020//sm is meaning an address and *sm means the value
	fmt.Printf("%T %v\n", sm, sm)   // same address as al (0xc0001160e0)
	fmt.Printf("%T %v\n", &sm, &sm) //shows different address
}

func main() {
	// assume 77 has been assigned in x
	// x:=77
	// we can see the variable's address using &x and value using (*(&x))
	// for example, we can get the address of x variable, &x=0x0301 and value , *0x0301=77
	i := 10                       //10 has been kept in i
	fmt.Printf("%T %v\n", &i, &i) // here &i returns memory address where variable i is.

	//*int 0xc000126010 here, *int means that it's a pointer and the value kept on the variable is an integer value (10), and the address is 0xc000126010
	fmt.Printf("%T %v\n", *(&i), *(&i)) //*(&i) prints the value on the address
	//int 10; here the value kept in i is 10 and the type is integer type

	//pointers
	//var <pointerName> *<data_type>
	//var ptr_1 *int

	//initialize pointer
	//var <pointerName> *<data_type> = &<variable_name>
	// i:=10
	// var ptr_1 *int = &i

	s := "Hello"
	//&s returns the address of s
	var b *string = &s //assigning the value of s to the pointer b
	fmt.Println(b)

	var a = &s //var a gets the address of s
	fmt.Println(a)

	c := &s // var c gets the address of c
	fmt.Println(c)

	//dereferencing a pointer
	*c = "Hello world" //changed the value in the address of c which is actually the address of b,a and s

	//b,a,c is an address and s is a variable
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*a)
	fmt.Println(s)

	//passing by value
	fmt.Println("Passing by value-------")
	ak := "Hello" //saved in address 0xc0000140c0
	fmt.Println(ak)
	fmt.Println(&ak)
	fmt.Println("-------")
	modify_by_value(ak) //using modify function, check the address
	fmt.Println("-------")
	fmt.Printf("%T %v\n", ak, ak)
	fmt.Printf("%T %v\n", &ak, &ak)

	//passing by reference
	fmt.Println("Passing by reference-------")
	al := "Hello1" //saved in address 0xc0001160e0
	fmt.Println(al)
	fmt.Println(&al) //0xc0001160e0
	fmt.Println("-------")
	modify_by_reference(&al) //using modify function, check the address
	fmt.Println("-------")
	fmt.Printf("%T %v\n", al, al)
	fmt.Printf("%T %v\n", &al, &al) //0xc0001160e0

}
