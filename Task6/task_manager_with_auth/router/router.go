package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/tasks", middleware.AuthMiddleware(false), controllers.GetTasks)
    r.GET("/tasks/:id", middleware.AuthMiddleware(false), controllers.GetTaskByID)
    r.POST("/tasks", middleware.AuthMiddleware(true), controllers.CreateTask)
    r.PUT("/tasks/:id",middleware.AuthMiddleware(true), controllers.UpdateTask)
    r.DELETE("/tasks/:id",middleware.AuthMiddleware(true), controllers.DeleteTask)
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.POST("/promote/:username", middleware.AuthMiddleware(true), controllers.PromoteUser)
    return r
}
