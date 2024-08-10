# Task Management API Documentation


### [postman documentation](https://documenter.getpostman.com/view/37574343/2sA3s3Gr82)
## Introduction

The Task Management API provides endpoints for managing tasks, including creating, reading, updating, and deleting tasks. This API is built using the Go programming language and the Gin framework. Tasks are now stored in a MongoDB database.

## MongoDB Integration

### Prerequisites

1. **MongoDB Installation:**
   - Ensure MongoDB is installed and running. Download from the [official MongoDB website](https://www.mongodb.com/try/download/community) if not already installed.

2. **Go MongoDB Driver:**
   - Install the Go MongoDB driver:
     ```sh
     go get go.mongodb.org/mongo-driver/mongo
     ```

### Configuration

1. **Database Connection:**
   - Configure MongoDB connection in the `InitDatabase` function:
     ```go
     func InitDatabase() error {
         clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
         client, err := mongo.Connect(context.TODO(), clientOptions)
         if err != nil {
             return err
         }
     
         err = client.Ping(context.TODO(), nil)
         if err != nil {
             return err
         }
     
         log.Println("Connected to MongoDB!")
     
         collection = client.Database("taskdb").Collection("tasks")
         return nil
     }
     ```

2. **Database and Collection:**
   - The database is named `taskdb` and the collection is named `tasks`.

## Endpoints

### GET /tasks
- **Description**: Get a list of all tasks.
- **Response**:
    ```json
    [
        {
            "id": "60d21b4667d0d8992e610c85",
            "title": "Task 1",
            "description": "First task",
            "due_date": "2024-08-07T12:00:00Z",
            "status": "Pending"
        },
        {
            "id": "60d21b4667d0d8992e610c86",
            "title": "Task 2",
            "description": "Second task",
            "due_date": "2024-08-08T12:00:00Z",
            "status": "In Progress"
        },
        {
            "id": "60d21b4667d0d8992e610c87",
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
        "id": "60d21b4667d0d8992e610c85",
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
        "id": "60d21b4667d0d8992e610c88",
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
        "id": "60d21b4667d0d8992e610c85",
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

This API provides basic CRUD functionality for managing tasks. It integrates with MongoDB for data persistence. Future enhancements may include additional features and optimizations.
