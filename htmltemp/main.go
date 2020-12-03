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

var noteStore = make(map[string]Note)

var id int = 0
