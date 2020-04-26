package main

import (
	"io"
	"log"
	"main/dbconnection"
	"net/http"
)

func main() {
	// Hello world, the web server
	dbconnection.GetPostingDb()
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
