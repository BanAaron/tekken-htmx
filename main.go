package main

import (
	"github.com/a-h/templ"
	"net/http"
)

func main() {
	greeter := hello([]string{"Aaron", "Chris", "Drew"})

	mux := http.NewServeMux()
	mux.Handle("/", templ.Handler(greeter))
	_ = http.ListenAndServe(":3333", mux)
}
