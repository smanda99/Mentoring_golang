package main

import (
	"encoding/json"
	"fmt"
)

// type Student struct {
// 	Name    string   `json:"name"`
// 	Age     int      `json:"age"`
// 	Address []string `json:"address"`
// 	Phone int
// }

func main() {
	DecodingJsonData()
}

// Aim: JSON to text
// --> Json
// --> varibele

func DecodingJsonData() {

	data := []byte(
		`{
			"name":"supraja",
			"age":20,
			"address": ["hyd","TS"]
			"phone":12652653476
		}`)
	// fmt.Println(json.Valid(data))
	// var student Student
	// fmt.Println(student)
	// json.Unmarshal(data, &student)
	// fmt.Println(student.Name)
	// fmt.Println(student.Age)
	// fmt.Println(student.Address)

	var person1 map[string]interface{}
	json.Unmarshal(data, &person1)
	fmt.Println(person1)

	for key, value := range person1 {
		fmt.Printf(" %v value is %v \n", key, value)
	}

}
