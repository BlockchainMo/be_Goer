package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("template_3.html")

		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}
		/* version 1 */
		/*
			score := r.FormValue("score")
			num, _ := strconv.Atoi(score)
			err = tmpl.Execute(w, num)
			/*

			/* version 2 */

		err = tmpl.Execute(w, r)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Print("Starting Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))

}
