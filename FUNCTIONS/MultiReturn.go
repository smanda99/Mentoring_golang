package main

import "fmt"

func main() {
	// fmt.Println("Multiple Return values in Functions")
	// // fmt.Println(Hello("supraja", "B.tech"))
	// // res := Hello()

	// fmt.Println(23 / 0)

	// k, l, m := Hello("python", "java")
	// fmt.Println(k)
	// fmt.Println(l)
	// fmt.Println(m)

	// fmt.Println(Add(23, 45, 23, 56, 21, 90, 56))

	// num := []int{23, 45, 23, 56, 21, 90, 56}
	// res := 0
	// // res = 9+ 23
	// for i := 0; i < len(num); i++ {
	// 	res = res + num[i]
	// }

	num1 := 8
	num2 := 9

	fmt.Println(8 - 9)

	// diff := num1 - num2

	if num1 < num2 {
		diff := Sub(num1, num2)
		fmt.Println("This is -ve difference", diff)
	} else {
		diff := Sub(num1, num2)
		fmt.Println("This is +ve difference", diff)
	}

}

// "jshdgsj" "gdhs" jshdgsjgdhs and len(str)
// 				Input 					  Output
func Hello(str1 string, str2 string) (string, int, string) {
	return str1 + str2, len(str1 + str2), str2 + str1

}

// // num = {1,2,3,4,5,6,7,8,9,10} []

// num  := "dsdghg","dhgfh"

// 1 2 3 4 5 6

// 1+2+3+4+5+6

func Sub(n1, n2 int) int {
	return n1 - n2
}

// func Sub(num ...int) int {

// 	res := 0
// 	for i := 0; i < len(num); i++ {
// 		res = res - num[i]
// 	}
// 	return res

// }
