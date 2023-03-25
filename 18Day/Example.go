package main

import "fmt"

func main() {
	fmt.Println("pointers")

	// var num int
	// var num *int // --> this pointer to memory location
	// *num = 24
	// fmt.Println(num)

	n := 24
	n1 := &n
	fmt.Println(n1)
	fmt.Println(&n)

	// n = n * 10
	// fmt.Println(*n1)
	// fmt.Println(n)

	Example(&n)
	fmt.Println(n)
	fmt.Println(*n1)

}

func Example(num *int) {
	*num = *num * 16

}
