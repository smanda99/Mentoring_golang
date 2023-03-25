package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var a int
	// var b int

	// var num int = a + b

	// fmt.Println("enter the value a")
	// fmt.Scan(&a)

	// fmt.Println("enter the value b")
	// fmt.Scan(&b)

	// var num int = a + b
	// fmt.Println(num)

	//  Bufio

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name)
	}

	// arr = [12,1,2,1,2]a
	array, _ := reader.ReadSlice('\n')
	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		fmt.Println(string(array[i]))

	}

	// fmt.Println(string(13))

}

// {
// 	name: "supraja"
// 	roll: "1290"
// }
