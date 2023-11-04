package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hola", miHandler)
	http.ListenAndServe(":8080", nil)
}

func miHandler (res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hola")
}