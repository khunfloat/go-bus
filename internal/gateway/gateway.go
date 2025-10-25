package gateway

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var publicPaths = []string{
	"/auth/login",
	"/auth/register",
	"/exam",
}

var jwtSecret = []byte("your-secret-key")

func ApiGateway(app *fiber.App) fiber.Router {
	router := app.Group("/", jwtMiddleware)
	return router
}

func jwtMiddleware(c *fiber.Ctx) error {
	path := c.Path()

	if isPublic(path) {
		return c.Next()
	}

	// extract Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// verify token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	// put claims into context
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("claims", claims)

	return c.Next()
}

func isPublic(path string) bool {
	for _, p := range publicPaths {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}
