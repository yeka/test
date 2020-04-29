package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello world\n")
		fmt.Fprintln(w, "Host:", req.Host)
		fmt.Fprintln(w, "Request URI:", req.RequestURI)
		if len(req.Header) > 0 {
			fmt.Fprintln(w, "Headers:")
			for k, vx := range req.Header {
				for _, v := range vx {
					fmt.Fprintln(w, "  ", k, ":", v)
				}
			}
		}
		if req.Body != nil {
			b, _ := ioutil.ReadAll(req.Body)
			fmt.Fprintln(w, "\nBody:\n"+string(b))
		}
	})
	err := http.ListenAndServe(":8001", m)
	if err != nil {
		log.Println(err)
	}
}
