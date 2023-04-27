package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Working with routes of a server")
	// GetHomeRequest()
	// GetRequest()
	// PostDataRequest()
	PostFormRequest()

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

func PostFormRequest() {
	URL := "http://localhost:1111/postform"

	data := url.Values{}
	data.Add("name", "supraja")
	data.Add("rollno", "1290")
	data.Add("sub", "go")

	response, err := http.PostForm(URL, data)
	if err != nil {
		panic(err)

	}

	res, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(res))
}
