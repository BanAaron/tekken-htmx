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

	err := godump.Dump(aaron)
	if err != nil {
		fmt.Println(err)
	}

	hm := homepage()
	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(hm))
	err = http.ListenAndServe(":3333", mux)
	if err != nil {
		panic(err)
	}
}
