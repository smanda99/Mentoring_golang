package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//  comma oka syntax
	reader := bufio.NewReader(os.Stdin)

	val, err := reader.ReadString('\n') // "56" "st"

	// ans := int(val)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}
