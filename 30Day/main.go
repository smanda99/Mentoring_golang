package main

import (
	"fmt"
	"time"
)

func main() {
	go Great("Hello")
	time.Sleep(time.Second)
	go Great("Go")

}

func Great(s string) {
	for i := 0; i < 10; i++ {
		// time.Sleep(5 * time.Millisecond)
		fmt.Println(s)

	}

}
