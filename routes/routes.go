package routes

import (
	"natsas-place/controllers"
	"natsas-place/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// API v1 group
	v1 := app.Group("/api/v1")

	// Authentication routes (public)
	auth := v1.Group("/auth")
	{
		auth.Post("/register", controllers.Register)
		auth.Post("/login", controllers.Login)

	}

	// Protected routes (require authentication)
	// protected := v1.Group("/", middleware.AuthMiddleware())
	// {
	// 	// User profile routes
	// 	profile := protected.Group("/profile")
	// 	{
	// 		profile.Get("/", controllers.GetProfile)
	// 		profile.Put("/", controllers.UpdateProfile)
	// 		profile.Post("/change-password", controllers.ChangePassword)
	// 	}

	// 	// User info
	// 	protected.Get("/me", controllers.GetMe)
	// }

	// Admin routes (require admin privileges)
	// admin := protected.Group("/admin", middleware.AdminMiddleware())
	// {
	// 	// User management
	// 	adminUsers := admin.Group("/users")
	// 	{
	// 		adminUsers.Get("/", controllers.GetAllUsers)
	// 		adminUsers.Get("/:id", controllers.GetUserById)
	// 		adminUsers.Put("/:id", controllers.UpdateUser)
	// 		adminUsers.Delete("/:id", controllers.DeleteUser)
	// 	}
	// }

	// Public routes (no authentication required)
	// Will be implemented later:
	// - Properties listing
	// - Property details
	// - Amenities
	// - Search

	// Optional auth routes (authentication optional)
	// Will be implemented later:
	// - Properties with user context

	// Add rate limiting to sensitive endpoints
	rateLimited := v1.Group("/", middleware.RateLimitMiddleware(5, 60)) // 5 requests per minute
	{
		rateLimited.Post("/auth/register", controllers.Register)
		rateLimited.Post("/auth/login", controllers.Login)
		// rateLimited.Post("/auth/forgot-password", controllers.ForgotPassword)
		// rateLimited.Post("/auth/reset-password", controllers.ResetPassword)
	}
}
