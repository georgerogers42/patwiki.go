package main

import (
	"github.com/georgerogers42/pat.go"
	"net/http"
	"flag"
	"text/template"
)

func main() {
	port := flag.String("p", "5000", "Port to run on")
	flag.Parse()
	m := pat.New()
	m.Get("/*", http.HandlerFunc(hello))
	http.ListenAndServe(":"+*port, m)
}

var helloTpl = template.Must(template.ParseFiles("hello.html"))

func hello(w http.ResponseWriter, r *http.Request) {
	// Path variable names are in the URL.Query() and start with ':'.
	name := r.URL.Query().Get(":name")
	helloTpl.Execute(w, map[string]string{"name": name})
}
