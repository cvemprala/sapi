package sapi

import (
	"encoding/json"
	"net/http"
)

type CreateTodoHandler struct{}

func (h *CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var todo TodoItem
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todoList.AddTodoItem(&todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
