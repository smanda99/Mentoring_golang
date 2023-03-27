package main

import "fmt"

//  User defined data type
//  built in data type --> int,string,float..

type Student struct {
	Name   string
	Rollno int
	Marks  float32
}

type Employee struct {
	Name   string
	Salary float64
}

func main() {

	// arr = []

	fmt.Println("Methods in Golang")
	var st1 Student
	fmt.Println(st1)

	var st2 Student
	st2.Name = "supraja"
	st2.Rollno = 1290
	st2.Marks = 90.7

	fmt.Println(st2)

	st2.UpdateName()
	fmt.Println(st2)

	// var emp1 Employee
	// emp1.

}

func (s *Student) UpdateName() {

	s.Name = "python"

}

func (e *Employee) UpdateSalary() {
	e.Salary = 1200000

}
