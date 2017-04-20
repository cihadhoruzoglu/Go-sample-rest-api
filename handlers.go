package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"io/ioutil"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

}

func TodoShow(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	todoId := vars["todoId"]

	id := bson.ObjectIdHex(todoId)

	todo := RepoFindTodo(id)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}

}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}

func TodoRemove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	vars := mux.Vars(r)
	todoId := vars["todoId"]

	fmt.Println("Remove id", todoId)

	if !bson.IsObjectIdHex(todoId) {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)

	id := bson.ObjectIdHex(todoId)

	RepoDestroyTodo(id)
}
