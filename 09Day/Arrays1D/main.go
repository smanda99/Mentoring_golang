package main

import (
	"fmt"
)

func main() {

	var array [7]int // 0,1,2,3,4,5,6 if we are acceseing are assigning values greter than last index show error index out of bound
	// fmt.Println(array)

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	array[0] = 12
	array[1] = 89
	array[2] = 8
	array[3] = 23
	array[4] = 9
	array[5] = 8
	array[6] = 56

	fmt.Println(array)

	// sort.Sort(array)
	// [12 89 8 23 9 7 56]

	// var MultiArray [3][3]int

	// [
	//   [12,2,4]
	//   [1,2,4]
	//   [1,5,6]
	// ]

}
