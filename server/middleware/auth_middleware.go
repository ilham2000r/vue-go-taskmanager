package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtSecretKey := "SecretKey"
		authHeader := c.Get("Authorization")
		// check header (Bearer)
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Missing authorization header"})
		}
		// delete the "Bearer" in header keep JWT only
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		// verify token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", uint(claims["user_id"].(float64)))

		fmt.Println(claims["user_id"])

		return c.Next()
	}
}
