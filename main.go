package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Ivan")
}

func main() {
	fmt.Println("Hello, Kubernetes! I'm from Jenkins CI!")
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}
