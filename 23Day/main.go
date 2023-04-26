package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web modules")

	URL := "https://github.com/Manda-supraja26/Mentoring_golang"

	response, err := http.Get(URL)

	if err != nil {
		panic("error in getting the url response")
	}

	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

}
