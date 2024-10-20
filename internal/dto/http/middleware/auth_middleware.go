package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JWTProtected() fiber.Handler {
	// Use jwtware for Fiber JWT middleware
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("your-secret-key"),
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized access"})
}

// Middleware for role-based access control
func RequireRole(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from the context
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)

		// Check if the user has the required role
		if !strings.EqualFold(role, requiredRole) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden - insufficient privileges",
			})
		}
		return c.Next()
	}
}
