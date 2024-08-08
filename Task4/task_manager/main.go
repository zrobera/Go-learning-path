package main

import (
    "task_manager/router"
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080") // Run the server on port 8080
}
