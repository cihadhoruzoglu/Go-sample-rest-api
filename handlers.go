package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"

	"io/ioutil"

	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	session := getSession()

	// collection := session.DB("hockey").C("players").Find()

	var results Todos

	err := session.DB("hockey").C("players").Find(nil).All(&results)
	if err != nil {
		fmt.Println("Error occured")
	} else {
		fmt.Println("Results All: ", results)
	}

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

	id, err := strconv.Atoi(todoId)

	if err != nil {
		panic(err)
	}

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
	w.WriteHeader(http.StatusNoContent)

	vars := mux.Vars(r)
	todoId := vars["todoId"]

	id, err := strconv.Atoi(todoId)

	if err != nil {
		panic(err)
	}

	todo := RepoDestroyTodo(id)

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}
