package main

import "fmt"

func main() {
	fmt.Println("Switch cases")

	// n := 10

	// switch n {

	// case 1:
	// 	fmt.Println("it is 1")

	// case 11:
	// 	fmt.Println("it is 10")

	// case 19:
	// 	fmt.Print("it is 19")

	// default:
	// 	fmt.Println("no value matched")
	// }
	n := 9
	i := 1

	switch i {
	case 1:
		fmt.Println(n * 1)
		i++
		fallthrough
	case 2:
		fmt.Println(n * 2)
		i++
		fallthrough
	case 3:
		fmt.Println(n * 3)
		i++
		fallthrough

	case 4:
		fmt.Println(n * 4)
		i++
		fallthrough

	case 5:
		fmt.Println(n * 5)
		i++
		fallthrough
	case 6:
		fmt.Println(n * 6)
		i++
		fallthrough

	case 7:
		fmt.Println(n * 7)
		i++
		fallthrough
	case 8:
		fmt.Println(n * 8)
		i++
		fallthrough
	default:
		fmt.Printf("this is %v table", n)
	}

}
