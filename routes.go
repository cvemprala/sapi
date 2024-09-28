package sapi

import (
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/todos/", &GetTodoHandler{})
	mux.Handle("/todos", &CreateTodoHandler{})
	mux.Handle("/todos/", &UpdateTodoHandler{})
	mux.Handle("/todos/", &DeleteTodoHandler{})
	return mux
}
