package main

import (
	"testing"
	"time"
)

func TestNewTodoItem(t *testing.T) {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})

	if todo.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", todo.ID)
	}
	if todo.Title != "Buy groceries" {
		t.Errorf("Expected Title to be 'Buy groceries', got %s", todo.Title)
	}
	if todo.Description != "Milk, Bread, Eggs" {
		t.Errorf("Expected Description to be 'Milk, Bread, Eggs', got %s", todo.Description)
	}
	if todo.Completed {
		t.Errorf("Expected Completed to be false, got %v", todo.Completed)
	}
	if todo.Priority != 1 {
		t.Errorf("Expected Priority to be 1, got %d", todo.Priority)
	}
	if len(todo.Tags) != 2 || todo.Tags[0] != "shopping" || todo.Tags[1] != "errands" {
		t.Errorf("Expected Tags to be ['shopping', 'errands'], got %v", todo.Tags)
	}
	if len(todo.Subtasks) != 0 {
		t.Errorf("Expected Subtasks to be empty, got %v", todo.Subtasks)
	}
}

func TestUpdateTodoItem(t *testing.T) {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})

	newDueDate := time.Now().Add(72 * time.Hour)
	todo.Update("Buy groceries and more", "Milk, Bread, Eggs, and Butter", &newDueDate, 2, []string{"shopping", "errands", "important"})

	if todo.Title != "Buy groceries and more" {
		t.Errorf("Expected Title to be 'Buy groceries and more', got %s", todo.Title)
	}
	if todo.Description != "Milk, Bread, Eggs, and Butter" {
		t.Errorf("Expected Description to be 'Milk, Bread, Eggs, and Butter', got %s", todo.Description)
	}
	if todo.DueDate == nil || !todo.DueDate.Equal(newDueDate) {
		t.Errorf("Expected DueDate to be %v, got %v", newDueDate, todo.DueDate)
	}
	if todo.Priority != 2 {
		t.Errorf("Expected Priority to be 2, got %d", todo.Priority)
	}
	if len(todo.Tags) != 3 || todo.Tags[0] != "shopping" || todo.Tags[1] != "errands" || todo.Tags[2] != "important" {
		t.Errorf("Expected Tags to be ['shopping', 'errands', 'important'], got %v", todo.Tags)
	}
}

func TestDeleteTodoItem(t *testing.T) {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})

	todo.Delete()

	if todo.ID != 0 || todo.Title != "" || todo.Description != "" || todo.Completed || !todo.CreatedAt.IsZero() || !todo.UpdatedAt.IsZero() || todo.DueDate != nil || todo.Priority != 0 || len(todo.Tags) != 0 || len(todo.Subtasks) != 0 {
		t.Errorf("Expected TodoItem to be empty, got %v", todo)
	}
}

func TestMarkComplete(t *testing.T) {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})

	todo.MarkComplete()

	if !todo.Completed {
		t.Errorf("Expected Completed to be true, got %v", todo.Completed)
	}
}

func TestMarkIncomplete(t *testing.T) {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})

	todo.MarkComplete()
	todo.MarkIncomplete()

	if todo.Completed {
		t.Errorf("Expected Completed to be false, got %v", todo.Completed)
	}
}

func TestAddSubtask(t *testing.T) {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})

	subtask := NewTodoItem(2, "Buy milk", "", nil, 1, []string{"shopping"})
	todo.AddSubtask(*subtask)

	if len(todo.Subtasks) != 1 || todo.Subtasks[0].ID != 2 {
		t.Errorf("Expected Subtasks to contain subtask with ID 2, got %v", todo.Subtasks)
	}
}

func TestRemoveSubtask(t *testing.T) {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})

	subtask := NewTodoItem(2, "Buy milk", "", nil, 1, []string{"shopping"})
	todo.AddSubtask(*subtask)
	todo.RemoveSubtask(2)

	if len(todo.Subtasks) != 0 {
		t.Errorf("Expected Subtasks to be empty, got %v", todo.Subtasks)
	}
}

func TestAddTodoItemToList(t *testing.T) {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)

	if len(todoList.todos) != 1 {
		t.Errorf("Expected TodoList to contain 1 item, got %d", len(todoList.todos))
	}
	if todoList.todos[1].Title != "Buy groceries" {
		t.Errorf("Expected Title to be 'Buy groceries', got %s", todoList.todos[1].Title)
	}
}

func TestUpdateTodoItemInList(t *testing.T) {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)

	newDueDate := time.Now().Add(72 * time.Hour)
	todoList.UpdateTodoItem(1, "Buy groceries and more", "Milk, Bread, Eggs, and Butter", &newDueDate, 2, []string{"shopping", "errands", "important"})

	if todoList.todos[1].Title != "Buy groceries and more" {
		t.Errorf("Expected Title to be 'Buy groceries and more', got %s", todoList.todos[1].Title)
	}
	if todoList.todos[1].Description != "Milk, Bread, Eggs, and Butter" {
		t.Errorf("Expected Description to be 'Milk, Bread, Eggs, and Butter', got %s", todoList.todos[1].Description)
	}
	if todoList.todos[1].DueDate == nil || !todoList.todos[1].DueDate.Equal(newDueDate) {
		t.Errorf("Expected DueDate to be %v, got %v", newDueDate, todoList.todos[1].DueDate)
	}
	if todoList.todos[1].Priority != 2 {
		t.Errorf("Expected Priority to be 2, got %d", todoList.todos[1].Priority)
	}
	if len(todoList.todos[1].Tags) != 3 || todoList.todos[1].Tags[0] != "shopping" || todoList.todos[1].Tags[1] != "errands" || todoList.todos[1].Tags[2] != "important" {
		t.Errorf("Expected Tags to be ['shopping', 'errands', 'important'], got %v", todoList.todos[1].Tags)
	}
}

func TestDeleteTodoItemFromList(t *testing.T) {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)

	todoList.DeleteTodoItem(1)

	if len(todoList.todos) != 0 {
		t.Errorf("Expected TodoList to be empty, got %d", len(todoList.todos))
	}
}

func TestGetTodoItemFromList(t *testing.T) {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)

	retrievedTodo := todoList.GetTodoItem(1)

	if retrievedTodo == nil {
		t.Errorf("Expected to retrieve TodoItem with ID 1, got nil")
	}
	if retrievedTodo.Title != "Buy groceries" {
		t.Errorf("Expected Title to be 'Buy groceries', got %s", retrievedTodo.Title)
	}
}
