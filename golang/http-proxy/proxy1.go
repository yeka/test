package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	tr := &http.Transport{
		DisableCompression: true,
		DisableKeepAlives:  false,
	}
	client := &http.Client{Transport: tr}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h, b, err := send(r, client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		for k, vx := range h {
			for _, v := range vx {
				w.Header().Add(k, v)
			}
		}
		w.Write(b)
	})
	log.Fatal(http.ListenAndServe(":8002", nil))
}

func send(r *http.Request, client *http.Client) (h http.Header, b []byte, err error) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8001"+r.RequestURI, r.Body)
	if err != nil {
		return
	}
	req.Header = r.Header

	res, err := client.Do(req)
	if err != nil {
		return
	}
	h = res.Header
	defer res.Body.Close()
	if res.Body != nil {
		b, err = ioutil.ReadAll(res.Body)
	}
	return
}
