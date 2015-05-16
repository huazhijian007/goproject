package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HI, this is an example\n")
}

func main() {

	http.HandleFunc("/", handle)
	http.ListenAndServeTLS(":8080", "./server.crt", "./server.key", nil)
}
