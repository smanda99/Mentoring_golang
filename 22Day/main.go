package main

import (
	"bufio"
	"fmt"
	"os"
)

type Employee struct {
	Name        string
	ID          int
	Salary      float64
	Techstack   []string
	Project     string
	Performance float64
}

// key ,value
//  [python ,java,mysql,go] strco
//  employee --> [em1,em2,em3,em4,em5....]
// using maps
// emp1 : {Name,id,salary,Techstack,Project}
// emp2 : {name ,id ,salary,}
// emp3 :

func main() {
	Employes := map[string]Employee{}

	// Employes["emp1"] = Employee{Name: "supraja", ID: 1, Salary: 120000, Techstack: []string{"python", "Mysql", "Git"}, Performance: 55.9}
	// Employes["emp2"] = Employee{Name: "supraja", ID: 1, Salary: 120000, Techstack: []string{"python", "Mysql", "Git"}, Performance: 55.9}

	// fmt.Println(Employes)

	// fmt.Println(Employes["emp1"].Salary)
	var em3 Employee
	var em4 Employee
	Employes["emp3"] = UserInput(em3)

	fmt.Println(Employes)

	fmt.Println(em4)

}

func UserInput(em Employee) Employee {
	reader := bufio.NewReader(os.Stdin)
	Name, err := reader.ReadString('\n')
	// ID, err := reader.ReadString('\n')
	// Salary, err := reader.ReadString('\n')
	// Techstack, err := reader.ReadSlice('\n')
	// Project, err := reader.ReadString('\n')
	// // Performance
	if err != nil {
		fmt.Println(err)
	}

	em.Name = Name
	return em

}

// UpdateSalary Method to update the salary
func (e *Employee) UpdateSalary(sal float64) {
	e.Salary = sal
}

// UpdateProject Method to update
// the Project
func (e *Employee) UpdateProject(proj string) {
	e.Project = proj
}

// UpdatePerformance Method to update the Performance
func (e *Employee) UpdatePerformance(per float64) {
	e.Performance = per
}
