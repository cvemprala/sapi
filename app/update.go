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

	var updateRequest UpdateTodoRequest
	if err := validator.ValidateRequest(w, r, &updateRequest); err != nil {
		return
	}

	todoList.UpdateTodoItem(id, updateRequest.Title, updateRequest.Description, &updateRequest.DueDate, updateRequest.Priority, updateRequest.Tags)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(updateRequest)
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
