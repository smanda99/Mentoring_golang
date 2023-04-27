package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name     string `json:"name"`
	Rollno   int    `json:"rollno"`
	Subject  []string
	Marks    float64
	Password string `json:"-"`
}

// username

func main() {
	fmt.Println("dealing with Json Data")
	EncodeingJson()

}

func EncodeingJson() {

	student := []Student{
		{"supraja", 1290, []string{"python", "go"}, 90.5, "supraja"},
		{"bala", 1291, []string{"java", "go"}, 99, "bala@1290"},
		{"lala", 34, []string{"go"}, 77, "jsdjs"},
	}

	res, err := json.MarshalIndent(student, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	// fmt.Printf("%T", res[:1])

}
