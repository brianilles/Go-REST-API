package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(len(todos) + 1)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func main() {
	todos = append(todos, Todo{ID: "1", Task: "Make an API with Go", Completed: false}, Todo{ID: "2", Task: "Eat dinner", Completed: false}, Todo{ID: "3", Task: "Go to sleep", Completed: false})

	r := mux.NewRouter()

	r.HandleFunc("/api/todos", getTodos).Methods("GET")
	r.HandleFunc("/api/todos/{id}", getTodo).Methods("GET")
	r.HandleFunc("/api/todos", postTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", r))
}
