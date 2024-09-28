package sapi

import (
	"sync"
	"time"
)

type TodoItem struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DueDate     *time.Time
	Priority    int
	Tags        []string
	Subtasks    []TodoItem
}

type TodoList struct {
	todos map[int]*TodoItem
	mu    sync.RWMutex
}

func NewTodoItem(id int, title string, description string, dueDate *time.Time, priority int, tags []string) *TodoItem {
	todo := &TodoItem{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DueDate:     dueDate,
		Priority:    priority,
		Tags:        tags,
		Subtasks:    []TodoItem{},
	}
	return todo
}

func (t *TodoItem) Update(title string, description string, dueDate *time.Time, priority int, tags []string) {
	t.Title = title
	t.Description = description
	t.DueDate = dueDate
	t.Priority = priority
	t.Tags = tags
	t.UpdatedAt = time.Now()
}

func (t *TodoItem) Delete() {
	*t = TodoItem{}
}

func (t *TodoItem) MarkComplete() {
	t.Completed = true
	t.UpdatedAt = time.Now()
}

func (t *TodoItem) MarkIncomplete() {
	t.Completed = false
	t.UpdatedAt = time.Now()
}

func (t *TodoItem) AddSubtask(subtask TodoItem) {
	t.Subtasks = append(t.Subtasks, subtask)
	t.UpdatedAt = time.Now()
}

func (t *TodoItem) RemoveSubtask(subtaskID int) {
	for i, subtask := range t.Subtasks {
		if subtask.ID == subtaskID {
			t.Subtasks = append(t.Subtasks[:i], t.Subtasks[i+1:]...)
			t.UpdatedAt = time.Now()
			break
		}
	}
}

func NewTodoList() *TodoList {
	return &TodoList{
		todos: make(map[int]*TodoItem),
	}
}

func (tl *TodoList) AddTodoItem(todo *TodoItem) {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	tl.todos[todo.ID] = todo
}

func (tl *TodoList) UpdateTodoItem(id int, title string, description string, dueDate *time.Time, priority int, tags []string) {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	if todo, exists := tl.todos[id]; exists {
		todo.Update(title, description, dueDate, priority, tags)
	}
}

func (tl *TodoList) DeleteTodoItem(id int) {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	delete(tl.todos, id)
}

func (tl *TodoList) GetTodoItem(id int) *TodoItem {
	tl.mu.RLock()
	defer tl.mu.RUnlock()
	return tl.todos[id]
}
