package controllers

import (
	"net/http"
	"strconv"

	"natsas-place/middleware"
	"natsas-place/models"
	"natsas-place/repository"
	user_service "natsas-place/services/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GetAllUsers retorna todos os usuários (admin only)
func GetAllUsers(c *fiber.Ctx) error {
	users, err := repository.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch users",
		})
	}

	// Converter para UserResponse
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponse := models.UserResponse{
			ID:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Phone:           user.Phone,
			Role:            user.Role,
			IsVerified:      user.IsVerified,
			IsActive:        user.IsActive,
			ProfileImage:    user.ProfileImage,
			Gender:          user.Gender,
			Nationality:     user.Nationality,
			Language:        user.Language,
			EmailVerifiedAt: user.EmailVerifiedAt,
			LastLoginAt:     user.LastLoginAt,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   userResponses,
	})
}

// GetUserById retorna um usuário específico pelo ID (admin only)
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := user_service.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	// Converter para UserResponse
	userResponse := models.UserResponse{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Phone:           user.Phone,
		Role:            user.Role,
		IsVerified:      user.IsVerified,
		IsActive:        user.IsActive,
		ProfileImage:    user.ProfileImage,
		Gender:          user.Gender,
		Nationality:     user.Nationality,
		Language:        user.Language,
		EmailVerifiedAt: user.EmailVerifiedAt,
		LastLoginAt:     user.LastLoginAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   userResponse,
	})
}

// UpdateUser atualiza um usuário (admin only)
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var input models.UserUpdateInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	// Validação dos dados
	if err := validator.New().Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Validation failed",
			"details": err.Error(),
		})
	}

	// Buscar usuário
	user, err := user_service.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	// Atualizar campos
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Phone != "" {
		user.Phone = input.Phone
	}
	if input.Gender != "" {
		user.Gender = input.Gender
	}
	if input.Nationality != "" {
		user.Nationality = input.Nationality
	}
	if input.Language != "" {
		user.Language = input.Language
	}
	if input.ProfileImage != "" {
		user.ProfileImage = input.ProfileImage
	}

	// Salvar alterações
	user, err = user_service.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update user",
		})
	}

	userResponse := models.UserResponse{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Phone:           user.Phone,
		Role:            user.Role,
		IsVerified:      user.IsVerified,
		IsActive:        user.IsActive,
		ProfileImage:    user.ProfileImage,
		Gender:          user.Gender,
		Nationality:     user.Nationality,
		Language:        user.Language,
		EmailVerifiedAt: user.EmailVerifiedAt,
		LastLoginAt:     user.LastLoginAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    userResponse,
	})
}

// DeleteUser deleta um usuário (admin only)
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	currentUserID := middleware.GetCurrentUser(c)

	// Converter ID para uint para comparação
	targetID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid user ID",
		})
	}

	// Não permitir auto-deleção
	if uint(targetID) == currentUserID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Cannot delete your own account",
		})
	}

	// Verificar se o usuário existe
	_, err = user_service.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	// Deletar usuário
	err = user_service.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
