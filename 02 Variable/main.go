package main

import "fmt"

// var Name string
// var Rollno int
// var Marks float32

var (
	Name   string
	Rollno int
	Marks  float32
)
var n int = 18
var isLogged bool = true

func main() {
	// Name := "python"
	// Rollno = 19
	// Marks = 90.9
	// n := 10
	fmt.Println(n)
	// friend := "supraja"

	fmt.Println(Name)
	fmt.Println(Rollno)
	fmt.Println(Marks)
	Greet()
	Hello()
	read()
}

func Greet() {
	fmt.Println(n)
}

func Hello() {
	fmt.Println(isLogged)

}
func read() {
	Name := "supraja"
	Name = ""
	fmt.Println(Name)

}
