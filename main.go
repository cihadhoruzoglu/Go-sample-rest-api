package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getCollection() *mgo.Collection {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	c := s.DB("hockey").C("players")

	return c
}
