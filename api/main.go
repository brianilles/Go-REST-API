package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID        string `json:"id`
	Task      string `json:"task"`
	Completed bool   `json:completed`
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func main() {
	todos = append(todos, Todo{ID: "1", Task: "Make an API with Go", Completed: false}, Todo{ID: "2", Task: "Eat dinner", Completed: false}, Todo{ID: "3", Task: "Go to sleep", Completed: false})

	r := mux.NewRouter()

	r.HandleFunc("/api/todos", getTodos).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}
