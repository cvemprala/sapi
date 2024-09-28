# sapi

## TodoItem Data Structure

The `TodoItem` data structure is designed to represent a task in a todo list. It includes the following attributes:

- **ID**: A unique identifier for the todo item.
- **Title**: The title of the todo item.
- **Description**: An optional description of the todo item.
- **Completed**: A boolean indicating whether the todo item is completed.
- **CreatedAt**: The timestamp when the todo item was created.
- **UpdatedAt**: The timestamp when the todo item was last updated.
- **DueDate**: An optional due date for the todo item.
- **Priority**: The priority level of the todo item.
- **Tags**: A list of tags for categorizing the todo item.
- **Subtasks**: A list of optional subtasks (nested `TodoItem` instances).

## Example Usage

Here are some examples of how to create, update, and manage `TodoItem` instances:

### Creating a TodoItem

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	fmt.Println(todo)
}
```

### Updating a TodoItem

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	
	newDueDate := time.Now().Add(72 * time.Hour)
	todo.Update("Buy groceries and more", "Milk, Bread, Eggs, and Butter", &newDueDate, 2, []string{"shopping", "errands", "important"})
	fmt.Println(todo)
}
```

### Marking a TodoItem as Complete

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	
	todo.MarkComplete()
	fmt.Println(todo)
}
```

### Adding a Subtask to a TodoItem

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	
	subtask := NewTodoItem(2, "Buy milk", "", nil, 1, []string{"shopping"})
	todo.AddSubtask(*subtask)
	fmt.Println(todo)
}
```

### Removing a Subtask from a TodoItem

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	
	subtask := NewTodoItem(2, "Buy milk", "", nil, 1, []string{"shopping"})
	todo.AddSubtask(*subtask)
	
	todo.RemoveSubtask(2)
	fmt.Println(todo)
}
```

## TodoList Data Structure

The `TodoList` data structure is designed to manage a collection of `TodoItem` instances with thread-safe operations. It includes the following attributes and methods:

- **todos**: A map that holds the created `TodoItem` instances.
- **mu**: A read-write mutex to ensure thread-safe operations.

### Methods

- **AddTodoItem(todo *TodoItem)**: Adds a new `TodoItem` to the list.
- **UpdateTodoItem(id int, title string, description string, dueDate *time.Time, priority int, tags []string)**: Updates an existing `TodoItem` in the list.
- **DeleteTodoItem(id int)**: Deletes a `TodoItem` from the list.

### Example Usage

Here are some examples of how to create, update, and manage `TodoItem` instances within a `TodoList`:

### Creating a TodoList and Adding a TodoItem

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)
	fmt.Println(todoList)
}
```

### Updating a TodoItem in a TodoList

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)
	
	newDueDate := time.Now().Add(72 * time.Hour)
	todoList.UpdateTodoItem(1, "Buy groceries and more", "Milk, Bread, Eggs, and Butter", &newDueDate, 2, []string{"shopping", "errands", "important"})
	fmt.Println(todoList)
}
```

### Deleting a TodoItem from a TodoList

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	todoList := NewTodoList()
	dueDate := time.Now().Add(48 * time.Hour)
	todo := NewTodoItem(1, "Buy groceries", "Milk, Bread, Eggs", &dueDate, 1, []string{"shopping", "errands"})
	todoList.AddTodoItem(todo)
	
	todoList.DeleteTodoItem(1)
	fmt.Println(todoList)
}
```
