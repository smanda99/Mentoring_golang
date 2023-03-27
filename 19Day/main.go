//  19th Day 25-03-2023
// author {
// 	name = "supraja"
// 	email ="mandasupraja1365139@gmail.com"
// }
//  Topic: Methods

package main

type Student struct {
	Name   string
	rollno int
	marks  float64
}

func main() {

	// var num int
	// var st1 Student
	// st1.rollno = 1290
	// st1.Name = "supraja"
	// st1.marks = 98

	// st2 := Student{Name: "supraja", rollno: 1290, marks: 29}
	// // st2.

	// // fmt.Printf("%T", st2)
	// fmt.Println(&num)

	// // b1 := Branch{IT: "it",CSE: "cse",EEE: "eee",ECE: "ece"}
	// // b1.
}

func hello(n int) {

}

func (s *Student) updateRollno() {

	s.rollno = 89
	if s.Name == "" {
		s.Name = "person1"
	}

}

// func (b Branch) UpdateBranch() {

// }
