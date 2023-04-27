package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Url parsing")
	URL := "https://github.com/Manda-supraja26/Mentoring_golang?course=hdj"
	result, _ := url.Parse(URL)

	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)

}
