package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	r, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8123", nil)
	if err != nil {
		log.Println("Error creating request:", err)
		os.Exit(1)
	}
	for {
		log.Print("Calling...")
		res, err := c.Do(r)
		if err != nil {
			log.Println("Error:", err)
		}
		if res == nil || res.Body == nil {
			continue
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("Error reading:", err)
		} else {
			log.Println(string(b))
		}
	}
}
