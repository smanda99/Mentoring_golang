package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	channel := make(chan string, 3)

	go func() {
		channel <- "supraja"
		channel <- "Hepsi"
		channel <- "clara"
		close(channel)
	}()

	// data := <-channel

	// fmt.Println(data)
	// data = <-channel
	// fmt.Println(data)
	// data = <-channel
	// fmt.Println(data)

	for n := range channel {
		fmt.Println(n)
	}

	// close(channel)

}
