package main

import "fmt"

func main() {
	a := 67 // declaration
	a = 3   // Assingment
	//+=

	//  a = a+4
	a += 4 // Add assignment --> 7
	fmt.Println(a)

	// a= a-8
	a -= 8
	fmt.Println(a)

	// a= a*2
	a *= 2
	fmt.Println(a)

	// a = a/2
	a /= 2
	fmt.Println(a)

	// a = a% 2
	a = 9
	a %= 2
	fmt.Println(a)

	// 12 %9 -->2

}
