package main

import "fmt"

func main() {
	// fmt.Println("If else")

	// name := "supraja"
	// rollno := 90

	// if name == "sfdg" {
	// 	rollno = 70
	// 	fmt.Println("this is supraja")
	// }

	// name = "ghshgd"
	// fmt.Println(name)
	// fmt.Println(rollno)

	// example2 if statement

	n := 87

	if n%2 == 0 {
		// n++
		n = n + 1
		fmt.Println(n)
	}

	fmt.Println(n)

	// If else

	//  n --> even or n --> odd

	N := 12

	if N%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("odd")
	}

	// 10 -->

	// else if

	m := "supr"

	if m > "14" {
		fmt.Println("greater")
	} else if m < "14" {
		fmt.Println("less")
	} else {
		fmt.Println("equal")
	}

}
