package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title		string
	Description	string
	CreatedOn	time.Time
}

type EditNote struct {
	Note
	Id string
}

var noteStore = make(map[string]Note)

var id int = 0

func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{id}", editNote)
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "base", noteStore)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "base", nil)
}

func saveNote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	desc := r.PostFormValue("description")
	note := Note{title, desc, time.Now()}
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	http.Redirect(w, r, "/", 302)
}

func editNote(w http.ResponseWriter, r *http.Request) {
	var viewModel EditNote
	vars := mux.Vars(r)
	k := vars["id"]
	if note, ok := noteStore[k]; ok {
		viewModel = EditNote{note, k}
	}else {
		http.Error(w, "Could not find the resource to edit.", http.StatusBadRequest)
	}
	renderTemplate(w, "edit", "base", viewModel)
}