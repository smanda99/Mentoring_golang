package main

import (
	"fmt"
)

func main() {
	fmt.Println(Conc("supraja", "hello"))
	fmt.Println(Add(5.6, 8.9))

}

//  "supraja" "hello" --> "suprajaHello"

func Conc(str1 string, str2 string) string {

}

// 7.6 -->
func Add(f1 float64, f2 float64) float64 {
	return f1 + f2

}
