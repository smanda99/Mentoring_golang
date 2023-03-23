package main

import "fmt"

type Employee struct {
	Name   string
	ID     int
	Salary float64
}

func main() {

	// var name string
	var employee1 Employee
	fmt.Println(employee1)
	// name := "supraja","rollo"
	employee1.ID = 1290
	employee1.Name = "supraja"
	employee1.Salary = 7489271
	fmt.Println(employee1)

	employee1.PrintEmployee()

}

// dry --> a+b

// Methods --> data
// employee --> hey i'm name holding rollno and having salary -->

func (e Employee) PrintEmployee() {
	fmt.Printf("Hey this is %s holding rollno %v and having salary %v", e.Name, e.ID, e.Salary)

}

//  Recursion
//  Maps
//  Pointers
//  pointers function
//  Memory Management
