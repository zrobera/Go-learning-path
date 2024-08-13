package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"task_manager/Delivery/router"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase() (*mongo.Database, error) {

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	clientOptions := options.Client().ApplyURI(uri)
	
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	db := client.Database("taskdb")
	return db, nil
}

func main() {
	db, err := InitDatabase()
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
		return
	}

	// Define timeout duration
	timeout := time.Second * 10

	r := gin.Default()

	router.Setup(timeout, db, r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err.Error())
	}
}
