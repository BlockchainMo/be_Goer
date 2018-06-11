package main

import (
	"html/template"
	"log"
	"os"
)

type Package struct {
	Name     string
	NumFuncs int
	NumVars  int
}

func main() {
	//Create a template, New a name for it
	//using "{{}}"
	//. means root, the second parameter of Execute!
	tmpl, err := template.New("Go-Web").Parse(`
Package name: {{.Name}}
Number of functions: {{.NumFuncs}}
Number of Variables: {{.NumVars}}
`) //New a name for template

	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	//Using the template
	err = tmpl.Execute(os.Stdout, &Package{
		Name:     "go-web",
		NumFuncs: 12,
		NumVars:  1200,
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
