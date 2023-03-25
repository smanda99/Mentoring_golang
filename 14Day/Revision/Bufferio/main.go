package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// a := 9
	// b := 8

	// c := a+b

	// scanf , scan , scanln
	// var n string

	// // fmt.Scan(&n)
	// // fmt.Println(n)

	// fmt.Scanf("%d", &n)
	// fmt.Println(n)

	take := bufio.NewReader(os.Stdin)
	val, err := take.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(val)

}
