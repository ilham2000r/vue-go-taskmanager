package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ilham2000r/vue-go-taskmanagment/config"
	"github.com/ilham2000r/vue-go-taskmanagment/models"
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
	Priority    string `json:"priority" validate:"required,oneof=A B C D"`
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

	// convert dueDate from string to time.Time
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid date format. Use YYYY-MM-DD",
			"details": err.Error(),
		})

	}
	// check priority
	if req.Priority != "A" && req.Priority != "B" && req.Priority != "C" && req.Priority != "D" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid priority. Must be A, B, C, or D",
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

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskId, userId).First(&task).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	config.DB.Save(&task)
	return c.JSON(task)
}

func (tc *TaskController) UpdateStatus(c *fiber.Ctx) error {
	taskId := c.Params("id")
	userId := c.Locals("user_id").(uint)

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskId, userId).First(&task).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	status := struct {
		TaskStatus bool `json:"taskStatus"`
	}{}

	if err := c.BodyParser(&status); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	task.TaskStatus = status.TaskStatus
	config.DB.Save(&task)
	return c.JSON(task)
}

func (tc *TaskController) DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("id")
	userId := c.Locals("user_id").(uint)

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", taskId, userId).First(&task).Error; err != nil {
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
	result := config.DB.Where("id = ? AND user_id = ?", taskId, userId).First(&task)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":   "Task not found",
			"details": result.Error.Error(),
		})
	}

	return c.JSON(task)
}

func (tc *TaskController) SearchByFilter(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(uint)

	// get filter value from request body
	filter := new(TaskFilter)
	if err := c.BodyParser(filter); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid input format",
			"details": err.Error(),
		})
	}

	// query in db
	query := config.DB.Where("user_id = ?", userId)
	// condition status filter
	if filter.TaskStatus != nil {
		query = query.Where("task_status = ?", *filter.TaskStatus)
	}
	// condition priority filter
	if filter.Priority != "" {
		// check priority [A,B,C,D]
		if filter.Priority != "A" && filter.Priority != "B" &&
			filter.Priority != "C" && filter.Priority != "D" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid priority. Must be A, B, C, or D",
			})
		}
		query = query.Where("priority = ?", filter.Priority)
	}

	// get data
	var tasks []models.Task
	if err := query.Find(&tasks).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Could not retrieve tasks",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"tasks": tasks,
		"count": len(tasks),
	})
}
