package main

import (
	"net/http"

	"./pages"
)

func main() {
	http.HandleFunc("/view/", pages.View)
	http.HandleFunc("/edit/", pages.Edit)
	http.HandleFunc("/save/", pages.Save)

	http.ListenAndServe(":10002", nil)
}
