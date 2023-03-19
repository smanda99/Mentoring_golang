package main

import "fmt"

func main() {
	a := 12
	b := 90

	fmt.Println(a == b) // checks whether there a is having same value of b --> F
	fmt.Println(a != b) //  --> T
	// fmt.Println(a<>b ) // 0 1 2 3  4 5
	fmt.Println(a > b)  // a is greater than b --> F
	fmt.Println(a < b)  // a is lessthan b --> T
	fmt.Println(a >= b) // a is greater than and equals to b  --> F
	fmt.Println(a <= b) // a is lessthan than and equals to b -->T

}
