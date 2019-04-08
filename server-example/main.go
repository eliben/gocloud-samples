package main

import (
	"fmt"
	"net/http"

	"gocloud.dev/server"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello\n")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!\n")
	})

	s := server.New(nil)
	s.ListenAndServe("localhost:8080", mux)
}
