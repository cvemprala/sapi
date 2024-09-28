package sapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGetTodoHandler(t *testing.T) {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)

	req, err := http.NewRequest("GET", "/todos/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &GetTodoHandler{}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := json.Marshal(todo)
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}

func TestCreateTodoHandler(t *testing.T) {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})

	payload, _ := json.Marshal(todo)
	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := &CreateTodoHandler{}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	expected, _ := json.Marshal(todo)
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}

func TestUpdateTodoHandler(t *testing.T) {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)

	updatedTodo := *todo
	updatedTodo.Title = "Buy groceries and more"
	updatedTodo.Description = "Milk, Bread, Eggs, and Butter"
	updatedTodo.Priority = 2
	updatedTodo.Tags = []string{"shopping", "errands", "important"}

	payload, _ := json.Marshal(updatedTodo)
	req, err := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := &UpdateTodoHandler{}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected, _ := json.Marshal(updatedTodo)
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}

func TestDeleteTodoHandler(t *testing.T) {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)

	req, err := http.NewRequest("DELETE", "/todos/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := &DeleteTodoHandler{}
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}
