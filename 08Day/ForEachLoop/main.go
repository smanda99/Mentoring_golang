package main

import "fmt"

func main() {
	// fmt.Println("For each loop")

	var arr = []string{"supraja", "ds", "python", "java"}

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// index,val

	// index, val
	// 0,suprajaarr
	// 1, ds
	// 2 , python
	// 3,java -->

	// for each loop

	for _, val := range arr {
		fmt.Println()
		fmt.Println(val)

	}

	// while mimic using for loop in go

	//  n <= 20 and n = n//2 and n > 0

	n := 20

	for n > 0 && n <= 20 {
		n = n / 2
		fmt.Println(n)
	}

}
