package main

import "fmt"

func main() {

	var array = []string{"supraja", "hsdgj", "uuty", "hgewh", "hsdh"}
	fmt.Println(len(array))

	// append()

	array = append(array, "bob", "tom")
	fmt.Println(array)

	// length := len(array) + 2
	// fmt.Printf("%T", newarraylength)

	// here we are inseting "bob" and "tom" in front of the slice

	var array1 = []string{}

	array1 = append(array1, "bob", "tom")
	fmt.Println(array1)

	for i := 0; i < len(array); i++ {
		array1 = append(array1, array[i])
	}

	fmt.Println(array1)

	var array2 = []string{}

	array2 = append(array2, "bala", "mottu")
	fmt.Println(array2)

	for i := 0; i < len(array1); i++ {
		array2 = append(array2, array1[i])
	}

	fmt.Println(array2)

}
