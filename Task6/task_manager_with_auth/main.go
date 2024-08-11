package main

import (
	"fmt"
	"log"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	err := data.InitDatabase()
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
		return
	}

	fmt.Println("Connected to MongoDB!")

	r := router.SetupRouter()
	r.Run(":8080")
}
