package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func routes(r *mux.Router) *mux.Router {
	createTodo := CreateTodoHandler{}
	updateTodo := UpdateTodoHandler{}
	deleteTodo := DeleteTodoHandler{}
	getTodo := GetTodoHandler{}

	apiRouter := r.PathPrefix("/api/todos").Subrouter()
	apiRouter.Handle("", createTodo).Methods(http.MethodPost)
	apiRouter.Handle("/{id}", updateTodo).Methods(http.MethodPut)
	apiRouter.Handle("/{id}", deleteTodo).Methods(http.MethodDelete)
	apiRouter.Handle("/{id}", getTodo).Methods(http.MethodGet)
	return r
}
