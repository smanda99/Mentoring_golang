package main

import "fmt"

func main() {

	//  maps
	// key , value
	// arr =[1,273,27897,2386,28]

	// name  : supraja
	//  y : 28
	// z : 90
	// gfd : 78
	//  re : 89
	// 9:10

	// st1 st2 st3 st4 st5 st6

	// var students = {
	// 	stud1 : st1,
	// 	stud2 : st2
	// 	stud3 : st3
	// 	stud4 : st4
	// 	stud5 : st5
	// }

	// map --> inorder
	// array -> order

	// name := map[string]string

	// 1 2 9 a b c d e

	// array := []int{12, 3, 4, 67, 89, 90 ,90,80,80}

	// fmt.Println(&array[0])
	// fmt.Println(&array[1])
	// fmt.Println(&array[2])
	// fmt.Println(&array[3])

	var student = map[string]string{
		"st1": "supraja",
		"st2": "vinay",
		"st3": "llps",
	}

	// fmt.Println(student["st1"])

	//  add to the map

	student["st4"] = "bala"
	student["st5"] = "neha"
	// fmt.Println(student)

	// assigning value to existing element

	student["st3"] = "ramesh"
	fmt.Println(student)

	//  no repeated values in maps

}
