package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ola mundo")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server On")
	addr := ":1357"
	http.ListenAndServe(addr, nil)
}
