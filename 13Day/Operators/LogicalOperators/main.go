package main

import "fmt"

func main() {
	// Logical Operators

	// And --> &&
	a := 4
	b := 3
	// c := a &&b
	c := 4
	str := "shg"
	// F 			F         T
	fmt.Println(a == b && c == b || c == a)
	fmt.Println(a == b || a != a)
	if str != "" {
		fmt.Println("hello")
	}
}
