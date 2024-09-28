package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateTodoHandler struct{}

func (h CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var todo TodoItem
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todoList.AddTodoItem(&todo)
	w.Header().Set("Location", fmt.Sprintf("api/todos/%d", todo.ID))
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(todo)
}
