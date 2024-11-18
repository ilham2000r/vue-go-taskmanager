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

	api.Post("/tasks", taskController.CreateTask)               // create task
	api.Put("/tasks/:id", taskController.UpdateTask)            // Update task (name,description, due date, priority)
	api.Patch("/tasks/:id/status", taskController.UpdateStatus) // checked status (finished, in process)
	api.Delete("/tasks/:id", taskController.DeleteTask)         // delete task
	api.Get("/tasks", taskController.ListTasks)                 // list all task
	api.Get("/tasks/:id", taskController.GetTaskById)           // list task by id
	api.Post("/tasks/search", taskController.SearchByFilter)    // search task (status, priority)
}
