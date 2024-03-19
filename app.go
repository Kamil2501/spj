package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Działa")
	})

	http.ListenAndServe(":8080", nil)
}
