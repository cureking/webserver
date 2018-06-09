package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":80", nil)
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>this is my first go webserver!</h1>")
}
