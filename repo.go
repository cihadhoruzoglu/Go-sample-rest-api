package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

var (
	todos Todos
)

// Give us some used data

func init() {
	RepoSyncAllTodos()
}

func RepoSyncAllTodos() {
	collection := getCollection()

	err := collection.Find(nil).All(&todos)
	if err != nil {
		fmt.Println("Collection not found")
	}
}

func RepoFindTodo(id bson.ObjectId) Todo {

	u := Todo{}

	collection := getCollection()

	if err := collection.FindId(id).One(&u); err != nil {
		// return empty Todo if not found
		return u
	}

	return u
}

/*
* RepoCreateTodo ...
* returns created todo
 */
func RepoCreateTodo(t Todo) Todo {
	collection := getCollection()

	t.Id = bson.NewObjectId()

	collection.Insert(t)

	RepoSyncAllTodos()

	return t
}

func RepoDestroyTodo(id bson.ObjectId) {
	collection := getCollection()

	if err := collection.RemoveId(id); err != nil {
		fmt.Errorf("Could not find Todo with id of %d to delete", id)
	}

	RepoSyncAllTodos()
}
