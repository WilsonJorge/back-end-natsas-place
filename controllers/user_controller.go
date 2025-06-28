package controllers

import (
	"github.com/gofiber/fiber/v2"
	"natsas-place/models"
	userService "natsas-place/services/user"
	"natsas-place/utils"
	"net/http"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	bodyParseError := c.BodyParser(&user)

	if bodyParseError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": bodyParseError.Error(),
		})
	}

	userCreated, err := userService.CreateUserService(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": http.StatusCreated,
		"data":   userCreated,
	})

}

func GetAllUsers(c *fiber.Ctx) error {

	users, err := userService.GetAllUsersService()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := userService.GetUserById(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	var userData utils.UpdateUserInput

	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	//result,err:= ser

	return c.JSON(fiber.Map{
		"message": "Dados do usuário atualizados com sucesso",
		"data":    userData,
	})
}
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := userService.DeleteUser(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   "User deleted",
	})

}
