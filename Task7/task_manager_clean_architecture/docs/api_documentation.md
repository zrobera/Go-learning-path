# Task Management with clean architecture API Documentation

## Introduction

The Task Management API provides endpoints for managing tasks, including creating, reading, updating, and deleting tasks. This API is built using the Go programming language and the Gin framework. Tasks are now stored in a MongoDB database.

### [postman documentation](https://documenter.getpostman.com/view/37574343/2sA3s4nr9B)

## Environment Variables

Before running the application, make sure to set up the necessary environment variables. You can create a `.env` file in the project root directory and add the following variables:

```plaintext
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=your_jwt_secret_key
```

- **MONGO_URI**: The connection string for your MongoDB instance.
- **JWT_SECRET**: The secret key used for signing JWT tokens. Ensure this is a strong, unique key.

**Note:** Never push your `.env` file or any sensitive information to version control. You can add the `.env` file to your `.gitignore` to avoid accidentally committing it.

```plaintext
# .gitignore
.env
```

## MongoDB Integration

### Prerequisites

1. **MongoDB Installation:**
    
    - Ensure MongoDB is installed and running. Download from the [official MongoDB website](https://www.mongodb.com/try/download/community) if not already installed.
        
2. **Go MongoDB Driver:**
    
    - go get go.mongodb.org/mongo-driver/mongo
        

### Configuration

1. **Database Connection:**
    
    - func InitDatabase() error { clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") client, err := mongo.Connect(context.TODO(), clientOptions) if err != nil { return err } err = client.Ping(context.TODO(), nil) if err != nil { return err } log.Println("Connected to MongoDB!") collection = client.Database("taskdb").Collection("tasks") return nil }
        
2. **Database and Collection:**
    
    - The database is named `taskdb` and the collection is named `tasks`.
        

## Protected Endpoints

### Using the JWT Token

- **Authorization Header:**
    
    - All end points except login and register are protected.
        
    - For all protected endpoints, include the JWT token in the `Authorization` header:
        
        ``` http
          Authorization: Bearer jwt_token_here
        
         ```
        

### Note on Roles

- **Admin Role:**
    
    - Admin users have full access to all endpoints, including creating, updating, deleting tasks, and promoting other users.
        
- **User Role:**
    
    - Regular users can only access tasks and cannot perform administrative actions.

## Folder Structure

The project is organized into distinct folders, each representing a different layer of the architecture:

### Delivery

- **`main.go`**: Initializes the HTTP server, sets up routing, and configures dependencies.
    
- **`controllers/controller.go`**: Handles incoming HTTP requests and invokes appropriate use case methods.
    
- **`routers/router.go`**: Defines and initializes the routes for the API.
    

### Domain

- **`domain.go`**: Contains core business entities such as `Task` and `User` structs. Defines the data models and core logic used throughout the application.
    

### Infrastructure

- **`auth_middleWare.go`**: Implements middleware for handling authentication and authorization using JWT tokens.
    
- **`jwt_service.go`**: Provides functions for generating and validating JWT tokens.
    
- **`password_service.go`**: Includes functions for hashing and comparing passwords to ensure secure storage of user credentials.
    

### Repositories

- **`task_repository.go`**: Defines the interface and implementation for task data access operations.
    
- **`user_repository.go`**: Defines the interface and implementation for user data access operations.
    

### Usecases

- **`task_usecases.go`**: Implements business rules related to tasks, including creating, updating, retrieving, and deleting tasks.
    
- **`user_usecases.go`**: Implements business rules related to users, including registration, login, and user promotion.
    

## Design Decisions

- **Separation of Concerns**: The architecture is designed to separate different aspects of the application into distinct layers, making the codebase more modular and easier to maintain.
    
- **Dependency Injection**: Dependencies such as repositories, JWT services, and password hashing are injected into use cases and controllers. This approach enhances testability and flexibility.
    
- **Use of Interfaces**: Interfaces are utilized to define contracts for repositories and services, allowing for easier testing and the potential for swapping implementations without impacting business logic.
    

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

## Authentication Endpoints

### POST /register
- **Description**: Register a new user.
- **Request**:
    ```json
    {
        "username": "newuser",
        "password": "password123"
    }
    ```
- **Response**:
    ```json
    {
        "message": "User registered successfully!"
    }
    ```

### POST /login
- **Description**: Login an existing user.
- **Request**:
    ```json
    {
        "username": "existinguser",
        "password": "password123"
    }
    ```
- **Response**:
    ```json
    {
        "token": "jwt_token_here"
    }
    ```

### POST /promote
- **Description**: Promote a user to admin.
- **Request**:
    ```json
    {
        "username": "userToPromote"
    }
    ```
- **Response**:
    ```json
    {
        "message": "User promoted to admin!"
    }
    ```

## How to Use

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/zrobera/Go-learning-path.git
    ```
   
2. **Navigate to the Project Directory**:
    ```sh
    cd Go-learning-path/Task7/task_manager_with_clean_architecture
    ```

3. **Start the Server**: 
    ```sh
    go run main.go
    ```

4. **Test Endpoints**: Use Postman or curl to test the API endpoints. For example, to get all tasks:
    ```sh
    curl -X GET http://localhost:8080/tasks
    ```

5. **Expected Responses**: Each endpoint's response format is shown above. Ensure your requests match the expected format.

6. **Error Handling**: The API handles various error scenarios, returning appropriate HTTP status codes and messages.

## Conclusion

This API provides basic CRUD functionality for managing tasks. It integrates with MongoDB for data persistence and includes JWT-based authentication. Future enhancements may include additional features and optimizations.


