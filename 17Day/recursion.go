package main

import "fmt"

func main() {
	Example(1)

}

//  calling itself
func Example(n int) {
	//  base condition
	if n == 10 {
		return
	}

	fmt.Println(n)
	Example(n + 1)

}
