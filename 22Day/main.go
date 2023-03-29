package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	Name      string
	ID        int
	Salary    float64
	Techstack []string
	Project   string
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
	// var em3 Employee
	// var em4 Employee
	// Employes["emp3"] = UserInput(em3)
	var emp Employee

	for i := 0; i < 5; i++ {
		Employes["emp"+string(emp.ID)] = UserInput(emp)
		emp.ID += 1

	}
	fmt.Println(Employes)

}

func UserInput(em Employee) Employee {
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	salValue, _ := reader.ReadString('\n')
	techstack, err := reader.ReadSlice('\n')
	project, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	em.Name = name
	em.Salary, _ = strconv.ParseFloat(strings.TrimSpace(salValue), 64)
	for _, v := range techstack {
		em.Techstack = append(em.Techstack, string(v))
	}
	em.Project = project
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
