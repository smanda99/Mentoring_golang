package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	web := []string{
		"http://www.youtube.com",
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
	}
	for _, w := range web {
		go StatusCode(w)
		// time.Sleep(1 * time.Millisecond)
		wg.Add(1)
	}

	wg.Wait()

}

func StatusCode(url string) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%v the status code of %v \n", res.StatusCode, url)
	}
	wg.Done()

}
