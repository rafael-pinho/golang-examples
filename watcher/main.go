package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"./services"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/register", services.Register)

	log.Fatal(http.ListenAndServe(":10001", nil))
}
