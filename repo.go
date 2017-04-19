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
	collection := getCollection()

	err := collection.Find(nil).All(&todos)
	if err != nil {
		fmt.Println("Collection not found")
	}
}

func RepoFindTodo(id bson.ObjectId) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}

	// return empty Todo if not found
	return Todo{}
}

/*
* RepoCreateTodo ...
* returns created todo
 */
func RepoCreateTodo(t Todo) Todo {
	collection := getCollection()

	t.Id = bson.NewObjectId()

	collection.Insert(t)

	todos = append(todos, t)

	return t
}

func RepoDestroyTodo(id bson.ObjectId) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
