package main

import (
	"fmt"
	"strconv"
)

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
	myslice1 := make([]int, 5, 10) //slice [type, length, capacity]

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

	fmt.Println("------------------")
	//Appending slices
	myslice2 := []int{1, 2, 3}
	myslice3 := []int{10, 20, 30}
	myslice4 := append(myslice2, myslice3...)
	fmt.Printf("myslice4 = %v\n", myslice4)

	// Operators are quite similar to other
	// Conditions
	time := 20
	if time < 18 {
		fmt.Println("Good Day")
	} else {
		fmt.Println("Good Evening")
	}
	// Loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//Looping in arrays
	for idx, val := range cars {
		fmt.Printf("%v\t%v\n", idx, val)
	}
	// Or
	fruits := []string{"apple", "mango", "peach", "pear"}
	for _, fruit := range fruits {
		if fruit == "peach" {
			fmt.Println("I am peach")
		} else {
			fmt.Println("I am not peach")
		}
	}
	Message("Gerald")

	pew := addition(1, 2)
	fmt.Println(pew)

	// Get addition and the length of the number
	result, length := myFunc(20, 40)
	fmt.Printf("myFunc addition is %v and length is %v\n", result, length)

	// Call the struct
	var person1 Person
	person1.name = "Gerald"
	person1.age = 20

	fmt.Println("Name:", person1.name)
	fmt.Println("Age:", person1.age)

	// Maps
	var newmap = map[string]int{"one": 1, "two": 2}
	fmt.Printf("newmap\t%v\v\n", newmap)
}

// Person Struct
type Person struct {
	name string
	age  int
}

// Message Functions
func Message(fname string) {
	fmt.Printf("Hello, %s!\n", fname)
}

// Function returning value
func addition(a, b int) int {
	return a + b
}

// Function with multiple returns
func myFunc(x int, y int) (result int, length int) {
	result = x + y
	length = len(strconv.Itoa(result))
	return
}
