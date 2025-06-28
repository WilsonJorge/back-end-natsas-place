package routes

import (
	"github.com/gofiber/fiber/v2"
	"natsas-place/controllers"
)

func Setup(app *fiber.App) {

	v1 := app.Group("/api/v1")

	//users
	{
		v1.Post("users", controllers.CreateUser)
		v1.Get("users", controllers.GetAllUsers)
		v1.Get("users/:id", controllers.GetUserById)

	}

}
