package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilham2000r/vue-go-taskmanagment/controllers"
	"github.com/ilham2000r/vue-go-taskmanagment/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Auth routes
	auth := controllers.AuthController{}
	app.Post("/register", auth.CreateUser)
	app.Post("/login", auth.Login)

	// Task routes - all protected by auth middleware
	taskController := controllers.NewTaskController()
	api := app.Group("/api", middleware.AuthMiddleware())

	api.Post("/tasks", taskController.CreateTask)
	api.Put("/tasks/:id", taskController.UpdateTask)
	api.Patch("/tasks/:id/status", taskController.UpdateStatus)
	api.Delete("/tasks/:id", taskController.DeleteTask)
	api.Get("/tasks", taskController.ListTasks)
	api.Get("/tasks/:id", taskController.GetTaskById)
	api.Post("/tasks/search", taskController.SearchByFilter)
}
