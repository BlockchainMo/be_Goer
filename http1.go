package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("version 1"))
	})

	http.HandleFunc("/bye", sayBye)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye version 1"))
}
