package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true) //get collection

	c := session.DB("taskdb").C("categories")

	docM := map[string]string{
		"name": "Open Source",
		"description": "Tasks for open-source projects",
	}

	//insert a map object
	err = c.Insert(docM)
	if err != nil {
		log.Fatal(err)
	}

	docD := bson.D{
		{"name", "Project"},
		{"description", "Project Tasks"},
	}
	//insert a document slice
	var count int
	count, err = c.Count()
	err = c.Insert(docD)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d records inserted", count)
	}
}
