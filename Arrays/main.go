package main

import "fmt"

func main() {
	// Arrays

	// int
	//  array is also a variable but with sequence of elements
	var ArrayInt = [5]int{12, 23, 1, 2, 3}
	fmt.Println(ArrayInt)

	//  what if we are giving the elements beyond the number in []
	// var Array = [5]int{12, 23, 1, 2, 3,89} //error :index 5 is out of bounds (>= 5)
	// fmt.Println(Array)

	var Array = [5]int{12, 23, 1} //[12 23 1 0 0]
	fmt.Println(Array)

	// Access
	fmt.Println(Array[2])

	// Array = Array[1:4]
	// re assigneing the value
	Array[4] = 45
	fmt.Println(Array)

	// Arrays using String
	// st := "supraja" string is imutable
	// "" --> defualt  0""

	var ArrayString = [10]string{"supraja", "ramya", "tom", "l", "m", "n", "k"}
	fmt.Println(ArrayString)

	// len() -->

	fmt.Println(len(ArrayString))

	for i := 0; i < len(ArrayString); i++ {
		fmt.Println(ArrayString[i])
	}

}
