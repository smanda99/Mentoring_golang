package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Working with routes of a server")
	GetHomeRequest()
	GetRequest()
	PostDataRequest()

}

func GetHomeRequest() {
	URL := "http://localhost:1111/"

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))

}

func GetRequest() {
	URL := "http://localhost:1111/get"

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))
}

// `"shejedh`
func PostDataRequest() {
	URL := "http://localhost:1111/post"

	data := strings.NewReader(`
	{
		"name": "supraja",
		"rollno":1290
	}`)
	response, _ := http.Post(URL, "application/json", data)
	defer response.Body.Close()
	res, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(res))

}
