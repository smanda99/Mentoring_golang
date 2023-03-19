package main

import "fmt"

func main() {
	// Num := []int{1, 2, 3}
	// //fmt.Println("Number:", Num)
	// for i := 0; i < len(Num); i++ {
	// 	fmt.Println("Number:", Num[i])
	// }

	// append

	// num := []int{2, 3, 5, 6}

	// test := append(num, 9, 80, 89)

	// fmt.Println(test)

	//copy
	num := []int{2, 3, 5, 6, 78, 56, 43, 21}

	num1 := []int{7, 8, 9, 8, 90, 18}

	// num2 := append(num1,num)
	copy(num, num1)

	fmt.Println(num)

	fmt.Println((num,num1)

	// num = [2, 3, 5, 6, 78, 56, 43, 21]
	// num1 = [7, 8, 9, 8, 90, 18]

}

// append
//copy
//equal
//len
