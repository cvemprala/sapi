package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UpdateTodoHandler struct{}

func (h UpdateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todo TodoItem
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todoList.UpdateTodoItem(id, todo.Title, todo.Description, todo.DueDate, todo.Priority, todo.Tags)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(todo)
}

type DeleteTodoHandler struct{}

func (h DeleteTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todoList.DeleteTodoItem(id)
	w.WriteHeader(http.StatusNoContent)
}
