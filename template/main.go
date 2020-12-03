package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.New("test").Parse(`{{define "T"}}{{ $greet := "Hello" }}{{ $greet }}, {{.}}!
{{ $greet }} again. {{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "World")
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
