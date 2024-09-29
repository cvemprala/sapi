package main

type CreateTodoRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	Priority    int       `json:"priority" validate:"required"`
	Tags        []string  `json:"tags" validate:"required"`
}

type UpdateTodoRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	Priority    int       `json:"priority" validate:"required"`
	Tags        []string  `json:"tags" validate:"required"`
}
