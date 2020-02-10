package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v6"

	"github.com/yeka/test-elasticsearch/generator"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	s := ""
	d := generator.Generate(1000000)

	handlers := []Handler{Take1{}, Take2{}}
	for _, h := range handlers {
		fmt.Println("Working on:", h.Index())

		_, err := es.Indices.Delete([]string{h.Index()})
		if err != nil {
			fmt.Println("Delete index:", err)
			return
		}

		f, err := os.Open(h.Index() + "/index.json")
		if err != nil {
			fmt.Println("Open index.json:", err)
			return
		}
		res, err := es.Indices.Create(h.Index(), es.Indices.Create.WithBody(f))
		fmt.Println(res)
		if err != nil {
			fmt.Println("Put mapping:", err)
			return
		}
		if err := f.Close(); err != nil {
			fmt.Println("Closing index.json:", err)
			return
		}

		for i, v := range d {
			s += fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, v.ID, "\n")
			s += h.Marshal(v)
			if (i+1)%10000 == 0 {
				_, err := es.Bulk(strings.NewReader(s), es.Bulk.WithIndex(h.Index()), es.Bulk.WithDocumentType("_doc"))
				fmt.Println(err)
				//fmt.Println(res)
				s = ""
			}
		}
	}
}

type Handler interface {
	Index() string
	Marshal(generator.Data) string
}

type Take1 struct{}

func (Take1) Index() string { return "take1" }
func (Take1) Marshal(d generator.Data) string {
	return fmt.Sprintf(`{"id":%v,"attributes":[{"key":"Color","value":"%v"}, {"key":"Size","value":"%v"}]}%v`, d.ID, d.Attributes["Color"], d.Attributes["Size"], "\n")
}

type Take2 struct{}

func (Take2) Index() string { return "take2" }
func (Take2) Marshal(d generator.Data) string {
	return fmt.Sprintf(`{"id":%v,"attributes":["Color:%v", "Size:%v"]}%v`, d.ID, d.Attributes["Color"], d.Attributes["Size"], "\n")
}
