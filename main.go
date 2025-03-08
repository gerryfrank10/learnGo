package main

import "fmt"

const PI = 3.1415926

func main() {
	var a, b = 6, "Hello"
	var i = 15.5
	c, d := 5, "World"

	//int and float and string
	var x uint = 500
	var y float64 = 15.51
	txt3 := "Gerald"

	// Printing with Println
	fmt.Println(b, d, c, a)
	fmt.Println(PI)

	// Printing variables with Printf
	fmt.Printf("%v\n\n", i)
	fmt.Printf("%T\n", b)
	fmt.Printf("%#v\n", d)

	// Printing type and variables
	fmt.Printf("Type : %T, value: %v\n", x, x)
	fmt.Printf("Type : %T, value: %v\n", y, y)
	fmt.Printf("Type : %T, value: %v\n", txt3, txt3)

	// Arrays
	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]int{1, 3, 9}
	cars := [...]string{"Volvo", "BMW", "Mazda"} // Array indexing starts 0
	arr3 := [5]int{1: 10, 2: 40}

	// arrays and slice
	myslice1 := make([]int, 5, 10)

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(cars[0])
	fmt.Println(arr3)
	//Length of an array
	fmt.Println(len(arr3))

	fmt.Println("--------------")
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("Length = %d\n", len(myslice1))
	fmt.Printf("Cap = %d\n", cap(myslice1))
}
