package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"base/validator"
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
	if err := validator.ValidateRequest(w, r, &todo); err != nil {
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
