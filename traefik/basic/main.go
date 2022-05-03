package main

import (
	"fmt"
	"net/http"
)

var addr = "127.0.0.1:8100"

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world"))
	})

	fmt.Println("Server ready at", addr)
	http.ListenAndServe(addr, m)
}
