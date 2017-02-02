package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	//"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["todoId"])
	fmt.Println("in TodoShow len id=", len(todos), todoId)

	if (err == nil) && (todoId > 0) && (todoId <= len(todos)) {
		fmt.Fprintf(w, "Todo show: %v %T", todoId, todoId)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todos[todoId-1]); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

}


func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	//deadline := time.Now().Add(500*time.Millisecond)
 	//ch := make(chan Todo)
        fmt.Println("start TodoCreate ")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	//RepoCreateTodo(todo, ch)
	//t := <- ch
	fmt.Println("in TodoCreate ", t)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["todoId"])
	fmt.Fprintln(w, "TodoDelete:", len(todos), todoId)
	if (err == nil) && (todoId > 0) && (todoId <= len(todos)) {
		RepoDestroyTodo(todoId)
	}
}
