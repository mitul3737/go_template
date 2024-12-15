// Dynamically typed language is go
package main

import "fmt"

func main() {
	// String (16bytes)
	// variable : var <variable_name> <variable_type> =  <value>
	var name string = "Lisa" //string assigned to the variable name

	//integers are  divided into unit8 (1byte),unit16(2bytes),unit32,unit64, int8(1 byte),int16(2 bytes),int32,int64, int (4bytes or 8 bytes)
	var val int = 200  //unsigned integer
	var mal int = -900 //signed integer

	//floatare divided into float32(4 bytes) and float64(8 bytes)
	var f float64 = 80.09
	var fk float64 = 0.775

	//boolean (1byte)
	var b bool = true
	var bk bool = false

	fmt.Println(name)
	fmt.Println(val)
	fmt.Println(mal)
	fmt.Println(f)
	fmt.Println(fk)
	fmt.Println(b)
	fmt.Println(bk)

}
