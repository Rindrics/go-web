package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t, err := template.New("test").Parse(`{{define "T"}}{{ $greet := "Hello" }}{{ $greet }}, {{.}}!
{{ $greet }} again. {{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "<a>Hello</a>")
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
