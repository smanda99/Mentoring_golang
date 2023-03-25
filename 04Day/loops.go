package main

import "fmt"

func main() {

	// 0,1,2,3,4,5
	// *
	// * *
	// * * *
	// * * * *
	// 1*2 = 2

	//  1 way of doing loops
	// n := 10

	// for i := 1; i <= 10; i++ {

	// 	fmt.Printf(" %v * %v = %v \n", n, i, n*i)

	// }
	//  arr = [12,3,4,45,6]
	// 1 10

	//  for each
	i := 1

	for i < 100000 {
		fmt.Print(i)
		i = i % 2

	}

}

func hello() {

}
