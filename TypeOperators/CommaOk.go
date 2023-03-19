package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a = "vala"
	value, err := strconv.Atoi(a)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Print(value)
	}
}
