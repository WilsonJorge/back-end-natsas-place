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

	// Public routes (no authentication required)
	// Properties listing and search
	properties := v1.Group("/properties")
	{
		properties.Get("/", controllers.GetAllProperties)
		properties.Get("/available", controllers.GetAvailableProperties)
		properties.Get("/:id", controllers.GetPropertyById)
		properties.Get("/:id/bookings", controllers.GetPropertyBookings)
	}

	// Protected routes (require authentication)
	protected := v1.Group("/", middleware.AuthMiddleware())
	{
		// Property management (authenticated users)
		myProperties := protected.Group("/my-properties")
		{
			myProperties.Get("/", controllers.GetMyProperties)
			myProperties.Post("/", controllers.CreateProperty)
			myProperties.Put("/:id", controllers.UpdateProperty)
			myProperties.Delete("/:id", controllers.DeleteProperty)
			myProperties.Post("/:id/images", controllers.AddPropertyImage)
			myProperties.Delete("/images/:imageId", controllers.RemovePropertyImage)
		}

		// Property bookings
		bookings := protected.Group("/properties")
		{
			bookings.Post("/:id/book", controllers.BookProperty)
			bookings.Put("/bookings/:bookingId/status/:status", controllers.UpdateBookingStatus)
		}

		// Property ratings
		protected.Post("/properties/:id/rate/:rating", controllers.RateProperty)
	}

	// Admin routes (require admin privileges)
	admin := protected.Group("/admin", middleware.AdminMiddleware())
	{
		// User management
		adminUsers := admin.Group("/users")
		{
			adminUsers.Get("/", controllers.GetAllUsers)
			adminUsers.Get("/:id", controllers.GetUserById)
			adminUsers.Put("/:id", controllers.UpdateUser)
			adminUsers.Delete("/:id", controllers.DeleteUser)
		}

		// Property management (admin)
		adminProperties := admin.Group("/properties")
		{
			adminProperties.Get("/", controllers.GetAllProperties)
			adminProperties.Get("/:id", controllers.GetPropertyById)
			adminProperties.Put("/:id", controllers.UpdateProperty)
			adminProperties.Delete("/:id", controllers.DeleteProperty)
		}
	}

	// Add rate limiting to sensitive endpoints
	rateLimited := v1.Group("/", middleware.RateLimitMiddleware(5, 60)) // 5 requests per minute
	{
		rateLimited.Post("/auth/register", controllers.Register)
		rateLimited.Post("/auth/login", controllers.Login)
		// rateLimited.Post("/auth/forgot-password", controllers.ForgotPassword)
		// rateLimited.Post("/auth/reset-password", controllers.ResetPassword)
	}
}
