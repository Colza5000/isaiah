package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome, to your Go docker container ", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
