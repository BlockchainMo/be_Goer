package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello, v1 version"))
	//})

	mux.HandleFunc("/bye", sayBye)

	log.Println("Starting server")
	http.ListenAndServe(":8001", mux)
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, v3 version"))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, v3 version"))
}
