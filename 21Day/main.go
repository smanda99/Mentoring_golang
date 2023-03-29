package main

import "fmt"

func main() {

	// make --> fun

	// 1. slice
	// 2. map
	// 3. channel

	// var Arr = []int {}

	// var Arr = make(T,length)

	//  slice creation using Make
	var Arr = make([]int, 10)
	fmt.Println(Arr)
	Arr = append(Arr, 120, 190)
	fmt.Println(len(Arr))

	// map --> inorder
	//  key , value pairs in maps
	var Array = make(map[string]int, 2)
	fmt.Println(Array)

	Array["supraja"] = 1290
	Array["ram"] = 1289
	fmt.Println(Array)

}
