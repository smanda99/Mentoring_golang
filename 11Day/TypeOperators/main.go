package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int = 97 // "10"
	var M float32 = 19.89
	var c float32
	c = float32(N) + M
	fmt.Printf("%T", c)
	fmt.Println(c)

	var res string
	// res = string(N) //method 97 --> a

	// fmt.Println(res) // we are checking the result of string() method 10 --> symbol

	res = strconv.Itoa(N) // this step is for converting 10 --> "10"
	fmt.Println(res)
	fmt.Printf("%T", res)

	// var strconv

}
