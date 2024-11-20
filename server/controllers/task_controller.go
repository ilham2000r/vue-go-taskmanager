package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ilham2000r/vue-go-taskmanagment/config"
	"github.com/ilham2000r/vue-go-taskmanagment/models"
	"time"
)

// TaskController struct
type TaskController struct{}

// instance of TaskController
func NewTaskController() *TaskController {
	return &TaskController{}
}

// struct for get request
type CreateTaskRequest struct {
	TaskName    string `json:"taskName" validate:"required"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate" validate:"required"`
	Priority    string `json:"priority" validate:"required,oneof=high medium low"`
}

func (tc *TaskController) CreateTask(c *fiber.Ctx) error {
	// get data request
	req := new(CreateTaskRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid input format",
			"details": err.Error(),
		})
	}
	fmt.Println("data Received:", req)
	// convert dueDate from string to time.Time
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid date format. Use YYYY-MM-DD",
			"details": err.Error(),
		})

	}
	// check priority
	validPriorities := map[string]bool{"high": true, "medium": true, "low": true}
	if !validPriorities[req.Priority] {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid priority. Must be high, medium, low",
		})
	}

	task := &models.Task{
		TaskName:    req.TaskName,
		Description: req.Description,
		DueDate:     dueDate,
		Priority:    req.Priority,
		UserID:      c.Locals("user_id").(uint),
	}

	// create task in db by Create func from controllers/task_controller
	if err := config.DB.Create(&task).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not create task"})
	}

	return c.Status(201).JSON(task)
}

func (tc *TaskController) UpdateTask(c *fiber.Ctx) error {
	taskId := c.Params("id")
	userId := c.Locals("user_id").(uint)

	var existingTask models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskId, userId).Take(&existingTask).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	// struct for update
	var updateData struct {
		TaskName    string `json:"taskName"`
		Description string `json:"description"`
		DueDate     string `json:"dueDate"`
		Priority    string `json:"priority"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid input",
			"details": err.Error(),
		})
	}

	// convert dueDate from string ro time.time
	if updateData.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02", updateData.DueDate)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":   "Invalid date format",
				"details": err.Error(),
			})
		}
		existingTask.DueDate = dueDate
	}

	// update other
	existingTask.TaskName = updateData.TaskName
	existingTask.Description = updateData.Description
	existingTask.Priority = updateData.Priority

	if err := config.DB.Save(&existingTask).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update task",
			"details": err.Error(),
		})
	}

	return c.JSON(existingTask)
}

func (tc *TaskController) UpdateStatus(c *fiber.Ctx) error {
	taskId := c.Params("id")
	userId := c.Locals("user_id").(uint)

	// รับ status เป็น boolean
	var updateData struct {
		Status bool `json:"status"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid input",
			"details": err.Error(),
		})
	}

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskId, userId).First(&task).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	// อัพเดต boolean status
	task.TaskStatus = updateData.Status
	if err := config.DB.Save(&task).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update status"})
	}

	return c.JSON(task)
}

func (tc *TaskController) DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("id")
	userId := c.Locals("user_id").(uint)

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskId, userId).Take(&task).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	config.DB.Delete(&task)
	return c.SendStatus(204)
}

func (tc *TaskController) ListTasks(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)
	var tasks []models.Task

	config.DB.Where("user_id = ?", userId).Find(&tasks)
	return c.JSON(tasks)
}

// struct for task filter
type TaskFilter struct {
	TaskStatus *bool  `json:"taskStatus"` // ใช้ pointer เพื่อให้รับค่า null ได้
	Priority   string `json:"priority"`
}

func (tc *TaskController) GetTaskById(c *fiber.Ctx) error {
	taskId := c.Params("id")
	userId := c.Locals("user_id").(uint)

	var task models.Task
	result := config.DB.Where("id = ? AND user_id = ?", taskId, userId).Take(&task)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":   "Task not found",
			"details": result.Error.Error(),
		})
	}

	return c.JSON(task)
}

// func (tc *TaskController) SearchByFilter(c *fiber.Ctx) error {
// 	userId := c.Locals("user_id").(uint)

// 	// get filter value from request body
// 	filter := new(TaskFilter)
// 	if err := c.BodyParser(filter); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error":   "Invalid input format",
// 			"details": err.Error(),
// 		})
// 	}

// 	// query in db
// 	query := config.DB.Where("user_id = ?", userId)
// 	// condition status filter
// 	if filter.TaskStatus != nil {
// 		query = query.Where("task_status = ?", *filter.TaskStatus)
// 	}
// 	// condition priority filter
// 	if filter.Priority != "" {
// 		// check priority [A,B,C,D]
// 		if filter.Priority != "A" && filter.Priority != "B" &&
// 			filter.Priority != "C" && filter.Priority != "D" {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": "Invalid priority. Must be A, B, C, or D",
// 			})
// 		}
// 		query = query.Where("priority = ?", filter.Priority)
// 	}

// 	// get data
// 	var tasks []models.Task
// 	if err := query.Find(&tasks).Error; err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error":   "Could not retrieve tasks",
// 			"details": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"tasks": tasks,
// 		"count": len(tasks),
// 	})
// }
