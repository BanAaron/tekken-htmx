package main

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/yassinebenaid/godump"
	"net/http"
)

type Person struct {
	Id   int
	Name string
	Url  string
}

func NewPerson(id int, name string, url string) Person {
	return Person{
		Id:   id,
		Name: name,
		Url:  url,
	}
}

func main() {
	aaron := NewPerson(1, "Aaron", "https://google.co.uk")
	sarah := NewPerson(2, "Sarah", "https://facebook.com")

	_ = godump.Dump(aaron)
	fmt.Println(aaron, sarah)

	header := headerTemplate(aaron)
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(header))
	_ = http.ListenAndServe(":3333", mux)
}
