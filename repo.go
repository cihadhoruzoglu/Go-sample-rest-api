package main

import (
	"fmt"
)

var (
	currentID int
	todos     Todos
)

// Give us some used data

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host Meetup!"})
}

func RepoFindTodo(id int) Todo {
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
	currentID++
	t.Id = currentID
	todos = append(todos, t)

	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
