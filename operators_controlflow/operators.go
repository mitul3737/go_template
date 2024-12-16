package main

import "fmt"

func main() {
	var city string = "Dhaka"
	var city_2 string = "New York"

	//comparison
	fmt.Println(city == city_2)
	fmt.Println(city != city_2)
	fmt.Println(city >= city_2)
	fmt.Println(city <= city_2)

	//Arithmetic
	var a, b string = "foo", "bar"
	fmt.Println(a + b)

	var x, y float64 = 79.0, 75.66
	fmt.Printf("%.2f", x-y)
	fmt.Println(x * y)
	fmt.Println(x / y)

	var m, n int = 10, 5
	fmt.Println(m % n)
	m++
	fmt.Println(m)
	n--
	fmt.Println(n)

	//logical operators

	var k int = 10
	fmt.Println((k < 100) && (k < 200))
	fmt.Println((k < 100) || (k < 200))

	//Assignment operators
	m += n //m=m+n
	fmt.Println(m)
	m -= n // m=m-n
	fmt.Println(m)
	m %= n
	fmt.Println(m)

	//Bitwise operator
	var mo, no int = 12, 25
	//12 =    00001100
	//25=     00011001
	//12&25 = 00001000 = 8
	fmt.Println(mo & no)
	//12 =    00001100
	//25=     00011001
	//12|25 = 00011101 = 29
	fmt.Println(mo | no)
	//12 =    00001100
	//25=     00011001
	//12^25=  00010101  = 21 (when both bits are different, we get 1)
	fmt.Println(mo ^ no)

	//left-shift
	//212 = 11010100
	// 212<<1 (1 bit shift to left) = 110101000 = 424
	var xy int = 212
	z := xy << 1
	fmt.Println(z)

	//right shift
	//212 = 11010100
	// 212<<2 (2 bit shift to left) = 00110101 = 53
	var zx int = 212
	zk := zx >> 2
	fmt.Println(zk)

	//if-else
	var fruit string = "grasp"
	if fruit == "apple" {
		fmt.Println("Fruit is an apple")
	} else if fruit == "jackfruit" {
		fmt.Println("Fruit is Jackfruit")
	} else {
		fmt.Println("Fuit is not an apple")

	}

	//switch statement
	var sw int = 100
	switch sw {
	case 10:
		fmt.Println("i is 10")
	case 100, 200:
		fmt.Println("i is either 100 or 200")
	default:
		fmt.Println("i is neither 0,100 or 200")

	}

	var mw int = 100
	switch mw {
	case -4:
		fmt.Println("i is -4")
	case 100:
		fmt.Println("i is 100")
		fallthrough //once fallthrough comes, it makes sure that next case is executed
	case 20:
		fmt.Println("i is 20")
		fallthrough
	default:
		fmt.Println("i is janina")

	}

	//for loop
	//for initialization;condition;post
	for i := 1; i <= 3; i++ {
		fmt.Println(i * 1)
		if i == 2 {
			break
		}
	}

}
