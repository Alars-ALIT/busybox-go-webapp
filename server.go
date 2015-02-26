package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, I like %s!", r.URL.Path[1:])
	fmt.Println("Received request: s%", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Sttttarted app at :8080")
	http.ListenAndServe(":8080", nil)
}
