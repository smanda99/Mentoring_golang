package main

import "fmt"

func main() {
	N := 10

	var num = &N

	fmt.Println(num)
	fmt.Println(&N)
	// Variable
	Example(num)
	//  pointer value --> *
	fmt.Println(*num)
	fmt.Println(N)

}

func Example(n *int) {
	*n = *n * 100

}
