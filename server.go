package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ROBOT6, I fisti %s!", r.URL.Path[1:])
	fmt.Println("Received request: s%", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("ROBOT6 i  i i  gulli :8080")
	http.ListenAndServe(":8080", nil)
}
