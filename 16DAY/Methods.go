package main

import "fmt"

func main() {
	var student Student

	student.Name = "supraja"
	student.Gender = "Female"
	student.Id = 1290
	student.Grade = "c"
	student.Section = "B"
	student.Lang = "Telugu"

	student.ChangeGrade()
	fmt.Println(student)

	var student1 Student
	student1.Name = "Bala"
	student1.Id = 89
	student1.Grade = "B"
	student1.Lang = "english"
	student1.Section = "A"

	student1.ChangeGrade()

	student1.ChangeLang()
	fmt.Println(student1)

}

type Student struct {

	//  properties of the student
	Name    string
	Id      int
	Gender  string
	Grade   string
	Section string
	Lang    string
}

func (s1 *Student) ChangeGrade() {
	s1.Grade = "A+"
	// fmt.Println(s1.Grade)
}

func (s1 *Student) ChangeLang() {
	s1.Lang = "Python"
}

// Recurssion
// Pointer
// Maps
// Methods
//

//  Recursion --> a function calling itself
//  why we need the Recursion
//  Let's say we wanna print "hello world " -->5 times --> using functions
//  func print1(){ 1 and print2() }, fun print2(){2 ,print3()}, func print3(){3,print(4)}, func print4(4,print4()){} ,func print5(5,print5()){}
//  func print(n int):
// 		if n == 0{ return }
// 		print(n)
// 		print(n-1)

// pointers
//  why we are using pointers
//  when we normally pass the data into any function in that case we dont directly pass the actual memory instead we pass the copy of the function which create some irregularities.
//  var name *int --> this indication this is a pointer which store integer
//  &name --> memory location of the pointer refrencing
//
