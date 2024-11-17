package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ilham2000r/vue-go-taskmanagment/config"
	"github.com/ilham2000r/vue-go-taskmanagment/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthController struct{}

func (ac *AuthController) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not hash password"})
	}
	user.Password = string(hashedPassword)

	// Save user to db
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not create user"})
	}

	return c.Status(201).JSON(user)
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	jwtSecretKey := "SecretKey"
	loginData := new(models.User)
	if err := c.BodyParser(loginData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input"})
	}

	// Find user
	var user models.User
	if err := config.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found"})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid password"})
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 168).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": t})
}
