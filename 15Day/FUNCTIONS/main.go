package main

import "fmt"

func main() {

	// Calling a function
	// Add()
	// Add()
	// Add() 2+8
	// Add() 4+6
	// Add() 2+3
	add()
	fmt.Println("hello")
	Subb(7, 3)
	fmt.Println(Multi(3, 2)) // printing the return value from function
	k := Multi(2, 3) + 3     // storing value in a variable and appling some changes
	fmt.Println(k)

}

//  function declare
func add() {
	a := 12
	b := 13
	fmt.Println(a + b)

}

// parameters Sub function
//

func Subb(num1 int, num2 int) {
	fmt.Println(num1 - num2)

}

//

// 4 + 3

//  return the value

func Multi(num1 int, num2 int) int {
	return num1 * num2

}

//  multi return
// multi parameter
// recurssion

// todo

// //  addition
// //  subtration
// //  multiplication
// // division
// //  remainder

// 8-9
// num1 num2
// if num2 > num1 :

// 9-10 -1 --> it is negitive differce

// division
// 8/2
// 8/0

// remainder
// 8%0
// 8/8
