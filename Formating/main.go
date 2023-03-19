package main

import "fmt"

func main() {
	// Formating

	fmt.Printf("\n") // --> New line

	// --> %v
	// A := 12
	// fmt.Printf("this Number is %v", A)

	// for strings %s
	s := "Manda"
	f := "supraja"
	fmt.Printf("This is FirstName %s and Second Name %s \n", f, s)

	fmt.Println("hello")

	// %g --> for float values
	number := 23.7

	fmt.Printf("This is float value: %v \n", number)

	//Boolean Values --> True or False
	isSigned := true

	fmt.Printf("This person is Logedin the status is : %t \n", isSigned)

	fmt.Printf("This is a type of %T", isSigned)

}
