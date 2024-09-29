package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"base/validator"
)

type CreateTodoHandler struct{}

func (h CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var todo TodoItem
	if err := validator.ValidateRequest(w, r, &todo); err != nil {
		return
	}

	todoList.AddTodoItem(&todo)
	w.Header().Set("Location", fmt.Sprintf("api/todos/%d", todo.ID))
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(todo)
}
