package main

import (
	"log"
	"os"

	"github.com/aalsa/management_dashboard/internal/db"
	"github.com/aalsa/management_dashboard/internal/handlers"
	"github.com/aalsa/management_dashboard/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	loadEnv()

	database, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	employeeHandler := handlers.NewEmployeeHandler(database)
	projectHandler := handlers.NewProjectHandler(database)
	taskHandler := handlers.NewTaskHandler(database)
	timeLogHandler := handlers.NewTimeLogHandler(database)
	authHandler := handlers.NewAuthHandler(database)

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.GET("/me", middleware.AuthMiddleware(), authHandler.GetMe)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/employees", employeeHandler.GetAll)
		api.POST("/employees", employeeHandler.Create)
		api.GET("/employees/:id", employeeHandler.GetByID)
		api.PUT("/employees/:id", employeeHandler.Update)
		api.DELETE("/employees/:id", employeeHandler.Delete)
		api.GET("/employees/:id/hours", timeLogHandler.GetEmployeeHours)

		api.GET("/projects", projectHandler.GetAll)
		api.POST("/projects", projectHandler.Create)
		api.GET("/projects/:id", projectHandler.GetByID)
		api.PUT("/projects/:id", projectHandler.Update)
		api.DELETE("/projects/:id", projectHandler.Delete)

		api.GET("/tasks", taskHandler.GetAll)
		api.POST("/tasks", taskHandler.Create)
		api.GET("/tasks/:id", taskHandler.GetByID)
		api.PUT("/tasks/:id", taskHandler.Update)
		api.DELETE("/tasks/:id", taskHandler.Delete)
		api.GET("/tasks/:id/hours", timeLogHandler.GetTaskHours)

		api.GET("/time-logs", timeLogHandler.GetAll)
		api.POST("/time-logs", timeLogHandler.Create)
		api.GET("/time-logs/:id", timeLogHandler.GetByID)
		api.PUT("/time-logs/:id", timeLogHandler.Update)
		api.DELETE("/time-logs/:id", timeLogHandler.Delete)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}

func loadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		file, err := os.Open(".env")
		if err != nil {
			log.Println("Warning: .env file exists but could not be opened")
			return
		}
		defer file.Close()

		var lines []string
		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if n > 0 {
				lines = append(lines, string(buf[:n]))
			}
			if err != nil {
				break
			}
		}

		for _, line := range lines {
			for i := 0; i < len(line); i++ {
				if line[i] == '=' {
					key := line[:i]
					value := line[i+1:]
					if len(value) > 0 && value[len(value)-1] == '\n' {
						value = value[:len(value)-1]
					}
					if len(value) > 0 && value[len(value)-1] == '\r' {
						value = value[:len(value)-1]
					}
					os.Setenv(key, value)
					break
				}
			}
		}
	}
}
