package main

import "fmt"

func main() {

	// a , b
	// operations  ---> + ,-,*,%,/ ,

	// >= , <= , >,<,!=,++,--

	// -------  ">" greater

	var a = 10
	var b = 90

	fmt.Println(a > b)

	//	// -------  "<" lessthan

	fmt.Println(a < b)

	//  ">="
	if a >= b {
		fmt.Println(" a is greater")
	} else {
		fmt.Println("b is greater")
	}

	// "<="

	if a <= b {
		fmt.Println("b is greater")
	} else {
		fmt.Println("a is greater")
	}

	// ==

	if a == b {
		fmt.Println("a is equals to b")
	} else {
		fmt.Println("Not equal")
	}

}
