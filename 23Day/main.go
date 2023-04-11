package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Http modules")
	url := "https://roadmap.sh/devops"
	reader, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer reader.Body.Close()

	data, err := ioutil.ReadAll(reader.Body)
	fmt.Println(string(data))

}
