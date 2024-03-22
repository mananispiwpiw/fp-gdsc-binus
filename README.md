# fp-gdsc-binus

This is a simple API for a task management application. It allows you to create, read, update, and delete tasks.

## Getting Started

1. Clone the repository
    ```
    git clone https://github.com/mananispiwpiw/fp-gdsc-binus.git
    ```
2. Navigate to the project's directory
    ```
    cd fp-gdsc-binus
    ```
3. Install the dependencies

    ```golang
    go mod tidy
    ```

4. Run the database migrations

    ```
    make updb
    ```

5. Start the server
    ```golang
    go run main.go
    ```

## DB Migration

1. **Install go-migrate**
   You can install **go-migrate** using the go get command:
    ```
    go get -u github.com/golang-migrate/migrate/cmd/migrate
    ```

## API Documentation

### Base URL

All API requests are made to:

http://localhost:8080

### Endpoints

### GET /tasks

Return a list of all tasks.

**Response**

```json
[
    {
        "ID": 1,
        "Title": "Example Task",
        "Description": "This is an example task",
        "CreatedAt": "2022-01-01T00:00:00Z",
        "UpdatedAt": "2022-01-01T00:00:00Z"
    }
    // ...
]
```

### POST /tasks

Adds a new task.

**Request**

```json
{
    "Title": "Example Task",
    "Description": "This is an example task"
}
```

**Response**

```json
{
    "ID": 1,
    "Title": "Example Task",
    "Description": "This is an example task",
    "CreatedAt": "2022-01-01T00:00:00Z",
    "UpdatedAt": "2022-01-01T00:00:00Z"
}
```

### PUT /tasks/{id}

Updates an existing task.

**Request**

```json
{
    "Title": "Updated Task",
    "Description": "This is an updated task"
}
```

**Response**

```json
{
    "message": "Task updated successfully!"
}
```

### DELETE /tasks/{id}

Deletes an existing task.

**Response**

```json
{
    "message": "Task deleted successfully!"
}
```

### Error Handling

If an error occurs, the API will return an HTTP status code and a JSON object with an _error_ property containing a description of the error. For example:

```json
{
    "error": "Invalid request method"
}
```

### Status Code

The API uses following status code:

-   **200 OK**: The request was successful.
-   **201 Created**: A new resource was successfully created.
-   **204 No Content**: The request was successful, but there's no representation to return (i.e. the response is intentionally empty).
-   **400 Bad Request**: The request could not be understood or was missing required parameters.
-   **404 Not Found**: A requested resource could not be found.
-   **500 Internal Server Error**: Something went wrong on the server.
