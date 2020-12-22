package main

import (
	"net/http"
)

func displayAnImage(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "cat.jpg")
}

func main() {
	http.HandleFunc("/" /*greet*/)
	http.ListenAndServe(":8080", nil)
}
