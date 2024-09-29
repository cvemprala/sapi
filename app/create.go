package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	s "sapi/base"
)

type CreateTodoHandler struct{}

func (h CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var createRequest CreateTodoRequest
	if err := s.ValidateRequest(w, r, &createRequest); err != nil {
		return
	}

	todo := TodoItem{
		ID:          len(todoList.todos) + 1,
		Title:       createRequest.Title,
		Description: createRequest.Description,
		DueDate:     &createRequest.DueDate,
		Priority:    createRequest.Priority,
		Tags:        createRequest.Tags,
	}

	todoList.AddTodoItem(&todo)
	w.Header().Set("Location", fmt.Sprintf("api/todos/%d", todo.ID))
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(todo)
}
