package main

import "fmt"

func main() {
	fmt.Println("Break Continue goto")

	// Break
	// i 0 to 10
	// i = 7-->
	// for i := 0; i < 10; i++ {
	// 	if i == 7 {

	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	// 7,8,9

	// }

	// continue

	// for i := 0; i < 10; i++ {
	// 	if i == 7 {
	// 		i := 10
	// i++
	// fmt.Println(i)

	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	// 7,8,9

	// }

	//  Go To

	for i := 0; i < 10; i++ {
		if i == 7 {
			goto lco
		}
		fmt.Println(i)

	}

lco:
	i := "supraja"

	fmt.Println(i)
}
