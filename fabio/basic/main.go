// This package provides http server used for testing purpose
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) == 1 {
		fmt.Println("Run: serve [addr] [text]")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(os.Args) > 2 {
			fmt.Fprintf(w, "<h1>%v</h1>", strings.Join(os.Args[2:], " "))
		} else {
			fmt.Fprint(w, "<h1>Hello world!</h1>")
		}
	})

	addr := ":80"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	log.Fatal(http.ListenAndServe(addr, nil))
}
