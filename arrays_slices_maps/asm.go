package main

import "fmt"

func main() {
	// arrays have fixed length, same type elements, capacity==length
	// var <array_name> [<size_of_array>] <data type>
	var grades [5]int
	fmt.Println(grades)

	//array initialization
	//var <arrayname> [<size>] <data type>= [<size of array>] <datatype> {vales}
	var fruits [2]string = [2]string{"apple", "orange"}
	fmt.Println(fruits)

	//array_name:=[<size of array>] <datatype> {vales}
	marks := [3]float64{23.4, 67.4, 29.9}
	fmt.Print(marks[1])

	for i := 0; i < len(marks); i++ {
		fmt.Println(grades[i])
	}

	//2d array
	//arr:=[row][col]{row contents1, row content 2,...}
	arr := [3][2]int{{2, 4}, {6, 4}, {8, 64}}
	fmt.Println(arr[2][1])

	//slice= not fixed, more flexible
	//slice_name:= []<data type> [values]
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	slice_1 := arr[0:1] //slice
	fmt.Println(slice_1)

	arr_1 := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_2 := arr_1[2:5]
	fmt.Println(slice_2)

	//declare and initialize slice
	//slice:=make([]<data_type,length,capacity)
	//slice:=make([]int,5,10)

	slice_3 := make([]int, 5, 8)
	fmt.Println(slice_3)      //default value is set to 0
	fmt.Println(cap(slice_3)) //capacity is 8
	fmt.Println(len(slice_3)) //length is 5 here

	//modify array
	slice_2[0] = 56909
	slice_2[1] = 3846743
	fmt.Println(arr_1) // value in arr_1 will change as slice_2's value have been changed which was referred to arr_1

	//apending to a slice
	//func append(slice of some datatype, values of same datatype)resulting are some array
	//slice = append(slice_1,10,20,30)
	slice_2 = append(slice_2, 900, 400, 500)
	fmt.Println(slice_2)
	fmt.Println(arr_1) //this impacted the main array once we added content to the slice, it replaced content on the referred array actually
	fmt.Println(len(arr_1))

	//adding 2 slice
	arr_3 := [5]int{1, 2, 4, 5, 3}
	slice_x := arr_3[0:2]
	slice_y := arr_3[2:]
	new_slice := append(slice_x, slice_y...) //we must add ... after  the second one
	fmt.Println(new_slice)

	//maps
	//<mapname>:= map[key_data_type]value_data_type{key-value-pairs}
	codes := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes) //map[en:English fr:French]

	codes_1 := make(map[string]int) //without initializing any value
	fmt.Println(codes_1)

	//value from map
	fmt.Println(codes["en"])
	fmt.Println(codes["fr"])

	//look for a key
	values, found := codes["en"] //found gives us value if we found the key or not
	fmt.Println(found, values)
	values_1, found_1 := codes["bn"]
	fmt.Println(found_1, values_1)

	//adding more value into the map
	codes["bn"] = "Bangla"
	fmt.Println(codes["bn"])
	for key, value := range codes {
		fmt.Println(key, "=>", value)
	}
}
