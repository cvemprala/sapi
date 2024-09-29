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

## HTTP Endpoints

The following HTTP endpoints are available for interacting with the todo items:

### GET /todos/{id}

Retrieve a todo item by its ID.

#### Request

- Method: GET
- URL: `/todos/{id}`

#### Response

- Status: 200 OK
- Body: JSON representation of the todo item

#### Example

```sh
curl -X GET http://localhost:8080/api/todos/1
```

### POST /todos

Create a new todo item.

#### Request

- Method: POST
- URL: `/todos`
- Body: JSON representation of the new todo item

#### Response

- Status: 201 Created
- Body: JSON representation of the created todo item

#### Example

```sh
curl -X POST http://localhost:8080/api/todos -d '{
  "ID": 1,
  "Title": "Buy groceries",
  "Description": "Milk, Bread, Eggs",
  "DueDate": "2023-12-31T23:59:59Z",
  "Priority": 1,
  "Tags": ["shopping", "errands"]
}' -H "Content-Type: application/json"
```

### PUT /todos/{id}

Update an existing todo item.

#### Request

- Method: PUT
- URL: `/todos/{id}`
- Body: JSON representation of the updated todo item

#### Response

- Status: 200 OK
- Body: JSON representation of the updated todo item

#### Example

```sh
curl -X PUT http://localhost:8080/api/todos/1 -d '{
  "Title": "Buy groceries and more",
  "Description": "Milk, Bread, Eggs, and Butter",
  "DueDate": "2024-01-01T23:59:59Z",
  "Priority": 2,
  "Tags": ["shopping", "errands", "important"]
}' -H "Content-Type: application/json"
```

### DELETE /todos/{id}

Delete a todo item by its ID.

#### Request

- Method: DELETE
- URL: `/todos/{id}`

#### Response

- Status: 204 No Content

#### Example

```sh
curl -X DELETE http://localhost:8080/api/todos/1
```

## Running Integration Tests

To run the integration tests, you need to have Python installed on your system. Follow the steps below:

1. Navigate to the root directory of the project.
2. Run the following command to execute all the integration tests:

```sh
pytest integration
```

This command will discover and run all the integration tests in the `integration` folder.

To run a single integration test, use the following command:

```sh
pytest integration/test_endpoints.py::test_create
```

Replace `test_create` with the name of the specific test method you want to run.
