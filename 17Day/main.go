package main

import "fmt"

func main() {
	// Recursion
	//  1 to 5 using function
	Fun1()
}

func Fun1() {
	fmt.Println(1)
	fun2()
}

func fun2() {
	fmt.Println(2)
	fun3()

}

func fun3() {
	fmt.Println(3)
	fun4()
}

func fun4() {
	fmt.Println(4)
	fun5()
}

func fun5() {
	fmt.Println(5)

}

//  1 to n
//  a function calling itself is called as recurssion
// 1 2

//  n = 10 --> re

func Printn(n int) {
	if n == 10 {
		fmt.Println(n)
		return
	}

	fmt.Println(n)
	Printn(n + 1)
}
