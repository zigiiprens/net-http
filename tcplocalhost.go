package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)

type myHandler struct{}

func ma(w http.ResponseWriter, r *http.Request) {
	//var pass string
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("ma.gtpl")
		t.Execute(w, nil)

	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	}
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
}

func main() {

	server := http.Server{
		Addr:    ":9090",
		Handler: &myHandler{},
	}


	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	mux["/"] = ma
	mux["/login"] = login
	server.ListenAndServe()
}
