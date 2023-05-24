package main

import (
	"fmt"
	"net/http"
)

// /hi

func hellofun(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		fmt.Fprintln(w, "the url path is wrong")
	}
	if r.Method != "GET" {
		fmt.Fprintln(w, "the Method is wrong")
	}
	w.Write([]byte("<h1>Hello This is Static Webserver</h1>"))

}

func formfun(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error parsing and the err: %v", err)
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	fmt.Fprintf(w, "the name : %v", name)
	fmt.Fprintf(w, "the password : %v", password)

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellofun)
	http.HandleFunc("/form", formfun)
	fmt.Println("the server is running on port 9000")
	http.ListenAndServe(":9000", nil)

}
