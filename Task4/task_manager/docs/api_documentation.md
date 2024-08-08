# Task Management API Documentation

## Introduction

The Task Management API provides endpoints for managing tasks, including creating, reading, updating, and deleting tasks. This API is built using the Go programming language and the Gin framework. The tasks are stored in an in-memory database for simplicity.

## Endpoints

### GET /tasks
- **Description**: Get a list of all tasks.
- **Response**:
    ```json
    [
        {
            "id": "1",
            "title": "Task 1",
            "description": "First task",
            "due_date": "2024-08-07T12:00:00Z",
            "status": "Pending"
        },
        {
            "id": "2",
            "title": "Task 2",
            "description": "Second task",
            "due_date": "2024-08-08T12:00:00Z",
            "status": "In Progress"
        },
        {
            "id": "3",
            "title": "Task 3",
            "description": "Third task",
            "due_date": "2024-08-09T12:00:00Z",
            "status": "Completed"
        }
    ]
    ```

### GET /tasks/:id
- **Description**: Get the details of a specific task.
- **Response**:
    ```json
    {
        "id": "1",
        "title": "Task 1",
        "description": "First task",
        "due_date": "2024-08-07T12:00:00Z",
        "status": "Pending"
    }
    ```

### POST /tasks
- **Description**: Create a new task.
- **Request**:
    ```json
    {
        "title": "New Task",
        "description": "Description of the new task",
        "due_date": "2024-08-08T12:00:00Z",
        "status": "Pending"
    }
    ```
- **Response**:
    ```json
    {
        "id": "1",
        "title": "New Task",
        "description": "Description of the new task",
        "due_date": "2024-08-08T12:00:00Z",
        "status": "Pending"
    }
    ```

### PUT /tasks/:id
- **Description**: Update a specific task.
- **Request**:
    ```json
    {
        "title": "Updated Task",
        "description": "Updated description",
        "due_date": "2024-08-09T12:00:00Z",
        "status": "In Progress"
    }
    ```
- **Response**:
    ```json
    {
        "id": "1",
        "title": "Updated Task",
        "description": "Updated description",
        "due_date": "2024-08-09T12:00:00Z",
        "status": "In Progress"
    }
    ```

### DELETE /tasks/:id
- **Description**: Delete a specific task.
- **Response**:
    ```json
    {
        "message": "Task deleted!"
    }
    ```

## How to Use

1. **Start the Server**: Ensure you have Go installed. Run the server using the following command:
    ```sh
    go run main.go
    ```

2. **Test Endpoints**: Use Postman or curl to test the API endpoints. For example, to get all tasks:
    ```sh
    curl -X GET http://localhost:8080/tasks
    ```

3. **Expected Responses**: Each endpoint's response format is shown above. Ensure your requests match the expected format.

4. **Error Handling**: The API handles various error scenarios, returning appropriate HTTP status codes and messages.

## Conclusion

This API provides basic CRUD functionality for managing tasks. Future enhancements may include integrating with a persistent database and adding more advanced features.

