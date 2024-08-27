package main

import (
	"io"
	"net/http"
)

func getMain(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	PORT := ":8080"

	http.HandleFunc("/", getMain)

	http.ListenAndServe(PORT, nil)
}
