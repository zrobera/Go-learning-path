package router

import (
	"task_manager/Delivery/controllers"
	infrastructure "task_manager/Infrastructure"
	repositories "task_manager/Repositories"
	usecases "task_manager/UseCases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	taskRouter := gin.Group("")
	NewTaskRouter(timeout, *db, taskRouter)

	userRouter := gin.Group("")
	NewUserRouter(timeout, *db, userRouter)
}


func NewTaskRouter(timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	tr := repositories.NewTaskRepository(db,"tasks")
	tc := &controllers.TaskController{
		TaskUseCase: usecases.NewTaskUseCase(tr, timeout),
	}

	jwtService := infrastructure.NewJWTService()

	authMiddleware := infrastructure.NewAuthMiddleware(jwtService)

	group.GET("/tasks",authMiddleware.AuthMiddleware(false), tc.GetTasks)
	group.GET("/tasks/:id", authMiddleware.AuthMiddleware(false), tc.GetTaskByID)
	group.POST("/tasks", authMiddleware.AuthMiddleware(true), tc.CreateTask)
	group.PUT("/tasks/:id", authMiddleware.AuthMiddleware(true), tc.UpdateTask)
	group.DELETE("/tasks/:id", authMiddleware.AuthMiddleware(true), tc.DeleteTask)
}

func NewUserRouter(timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	tr := repositories.NewUserRepository(db,"users")
	jwtService := infrastructure.NewJWTService()
	passwordService := infrastructure.NewPasswordService()

	tc := &controllers.UserController{
		UserUseCase: usecases.NewUserUseCase(tr, passwordService,jwtService,timeout),
	}

	authMiddleware := infrastructure.NewAuthMiddleware(jwtService)

	group.POST("/register", tc.Register)
	group.POST("/login", tc.Login)
	group.POST("/promote/:username", authMiddleware.AuthMiddleware(true), tc.PromoteUser)
}
