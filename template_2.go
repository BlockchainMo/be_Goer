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

		/* version 1, template is static */
		/*
					tmpl, err := template.New("Go-Web").Parse(`
			Package name: {{.Name}}
			Number of functions: {{.NumFuncs}}
			Number of Variables: {{.NumVars}}
					`) //New a name for template
		*/

		/* version 2, template parses from files*/
		tmpl, err := template.ParseFiles("template_2.html")

		if err != nil {
			fmt.Fprintf(w, "ParseFiles: %v", err)
			return
		}

		err = tmpl.Execute(w, &Package{
			Name:     "Go-Web",
			NumFuncs: 12,
			NumVars:  1200})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Print("Starting Server...")
	log.Fatal(http.ListenAndServe(":4000", nil))

}
