package main

import "fmt"

func main() {

	// var n *int

	num := 10
	var num1 = &num
	// fmt.Println(num1)
	// fmt.Println(&num)

	Example(num1)
	fmt.Println(*num1)

	// *num1 = *num1 * num

	// Example()

	// Example(num)
	fmt.Println(num)

}

func Example(num *int) {
	*num = *num * 18

}
