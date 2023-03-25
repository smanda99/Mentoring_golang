package main

import "fmt"

func main() {

	name := "GOLANG"

	fmt.Println(len(name)) //6

	for i := 1; i <= 3; i++ {
		fmt.Println(string(name[i]))
	}

}
