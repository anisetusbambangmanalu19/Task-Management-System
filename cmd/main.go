package main

import (
	"github.com/anisetusbambangmanalu19/task-management/internal/config"
	"github.com/anisetusbambangmanalu19/task-management/internal/delivery"
	"github.com/anisetusbambangmanalu19/task-management/internal/middleware"
	"github.com/anisetusbambangmanalu19/task-management/internal/repository"
	"github.com/anisetusbambangmanalu19/task-management/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend running 🚀",
		})
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	userRepo := repository.UserRepository{}
	userUsecase := usecase.UserUsecase{UserRepo: userRepo}
	userHandler := delivery.UserHandler{UserUsecase: userUsecase}
	projectRepo := repository.ProjectRepository{}
	projectUsecase := usecase.ProjectUsecase{ProjectRepo: projectRepo}
	projectHandler := delivery.ProjectHandler{ProjectUsecase: projectUsecase}
	taskRepo := repository.TaskRepository{}
	taskUsecase := usecase.TaskUsecase{TaskRepo: taskRepo, ProjectRepo: projectRepo}
	taskHandler := delivery.TaskHandler{TaskUsecase: taskUsecase}

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", userHandler.Profile)
	}

	auth.POST("/projects", projectHandler.Create)
	auth.GET("/projects", projectHandler.MyProjects)

	auth.POST("/projects/:project_id/tasks", taskHandler.Create)
	auth.GET("/projects/:project_id/tasks", taskHandler.GetByProject)

	r.Run(":8080")
}
