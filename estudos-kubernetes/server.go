package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HOME)
	http.ListenAndServe(":8080", nil)
}

func HOME(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Ola Mundo dos Kubernetes</h1>"))
}
