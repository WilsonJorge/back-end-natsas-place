package middleware

import (
	"strings"
	"time"

	"natsas-place/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware protege rotas que requerem autenticação
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Missing authorization header",
			})
		}

		// Verifica se o header começa com "Bearer "
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid authorization header format",
			})
		}

		tokenString := tokenParts[1]

		// Valida o token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid token",
			})
		}

		// Verifica se o token não está expirado
		if claims.ExpiresAt.Time.Before(time.Now()) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Token expired",
			})
		}

		// Adiciona as informações do usuário ao contexto
		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRole", claims.Role)

		return c.Next()
	}
}

// AdminMiddleware protege rotas que requerem privilégios de administrador
func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("userRole")

		if role != "ADMIN" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   true,
				"message": "Insufficient privileges",
			})
		}

		return c.Next()
	}
}

// OptionalAuthMiddleware permite autenticação opcional (não obrigatória)
func OptionalAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			// Se não há token, continua sem autenticação
			return c.Next()
		}

		// Verifica se o header começa com "Bearer "
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			// Token inválido, mas continua sem autenticação
			return c.Next()
		}

		tokenString := tokenParts[1]

		// Valida o token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			// Token inválido, mas continua sem autenticação
			return c.Next()
		}

		// Verifica se o token não está expirado
		if claims.ExpiresAt.Time.Before(time.Now()) {
			// Token expirado, mas continua sem autenticação
			return c.Next()
		}

		// Adiciona as informações do usuário ao contexto
		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRole", claims.Role)

		return c.Next()
	}
}

// GetCurrentUser retorna o ID do usuário atual do contexto
func GetCurrentUser(c *fiber.Ctx) uint {
	userID := c.Locals("userID")
	if userID == nil {
		return 0
	}

	if id, ok := userID.(uint); ok {
		return id
	}

	return 0
}

// GetCurrentUserEmail retorna o email do usuário atual do contexto
func GetCurrentUserEmail(c *fiber.Ctx) string {
	userEmail := c.Locals("userEmail")
	if userEmail == nil {
		return ""
	}

	if email, ok := userEmail.(string); ok {
		return email
	}

	return ""
}

// GetCurrentUserRole retorna o role do usuário atual do contexto
func GetCurrentUserRole(c *fiber.Ctx) string {
	userRole := c.Locals("userRole")
	if userRole == nil {
		return ""
	}

	if role, ok := userRole.(string); ok {
		return role
	}

	return ""
}

// IsAuthenticated verifica se o usuário está autenticado
func IsAuthenticated(c *fiber.Ctx) bool {
	return GetCurrentUser(c) > 0
}

// IsAdmin verifica se o usuário atual é admin
func IsAdmin(c *fiber.Ctx) bool {
	return GetCurrentUserRole(c) == "ADMIN"
}

// UserOwnsResource verifica se o usuário atual é dono do recurso
func UserOwnsResource(c *fiber.Ctx, resourceUserID uint) bool {
	currentUserID := GetCurrentUser(c)
	return currentUserID == resourceUserID || IsAdmin(c)
}

// RateLimitMiddleware implementa rate limiting básico
func RateLimitMiddleware(maxRequests int, windowSeconds int) fiber.Handler {
	// Armazena contadores de requisições por IP
	requests := make(map[string][]time.Time)

	return func(c *fiber.Ctx) error {
		clientIP := c.IP()
		now := time.Now()

		// Limpa requisições antigas
		if clientRequests, exists := requests[clientIP]; exists {
			var validRequests []time.Time
			for _, reqTime := range clientRequests {
				if now.Sub(reqTime).Seconds() < float64(windowSeconds) {
					validRequests = append(validRequests, reqTime)
				}
			}
			requests[clientIP] = validRequests
		}

		// Verifica se excedeu o limite
		if len(requests[clientIP]) >= maxRequests {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   true,
				"message": "Too many requests",
			})
		}

		// Adiciona a requisição atual
		requests[clientIP] = append(requests[clientIP], now)

		return c.Next()
	}
}
