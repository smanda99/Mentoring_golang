package main

import (
	"fmt"
)

func main() {
	fmt.Println("error")
	// var err error

	// name := "supraja"
	// if name == "supraja" {
	// 	panic("this is panic")

	// }

	// data, err := http.Get("https://github.com/Manda-supraja26/Mentoring_golang")
	// if err != nil {
	// 	panic("this is a panic situation")

	// }

	// fmt.Println("this is an application")
	// fmt.Println("this is end")

	//  when there is defer in the function when will panic get executed

	defer fmt.Println("this is an application")

	name := "supraja"
	if name == "supraja" {
		panic("this is panic")

	}

}
