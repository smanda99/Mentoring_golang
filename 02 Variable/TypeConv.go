package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// A := 12.0

	const pie float32 = 3.14159

	B := 97
	var c string

	// c = float32(A) + float32(B)
	// fmt.Printf("%T", c)

	c = strconv.Itoa(B)
	fmt.Println(c)
	fmt.Printf("%T", c)

	// c =
	//  "97"

	var st string
	// st =""
	fmt.Scanf(st)
	fmt.Println(st)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your output")
	N, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(N)

}
