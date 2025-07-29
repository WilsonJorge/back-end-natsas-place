package auth

import (
	"errors"
	"time"

	"natsas-place/database"
	"natsas-place/models"
	"natsas-place/utils"
)

func RegisterUser(input models.UserRegistrationInput) (*models.UserResponse, error) {

	var existingUser models.User
	result := database.Instance.Where("email = ?", input.Email).First(&existingUser)
	if result.Error == nil && existingUser.ID > 0 {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
		Role:     "USER",
		IsActive: true,
	}

	err = database.Instance.Create(&user).Error
	if err != nil {
		return nil, errors.New("failed to create user " + err.Error())
	}

	userResponse := &models.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		Role:         user.Role,
		IsVerified:   user.IsVerified,
		IsActive:     user.IsActive,
		ProfileImage: user.ProfileImage,
		LastLoginAt:  user.LastLoginAt,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return userResponse, nil
}

func LoginUser(input models.UserLoginInput) (*models.UserResponse, *utils.AuthResponse, error) {

	var user models.User
	result := database.Instance.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		return nil, nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, nil, errors.New("user account is inactive")
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return nil, nil, errors.New("invalid email or password, " + user.Password + " " + input.Password)
	}

	now := time.Now()
	user.LastLoginAt = &now
	database.Instance.Save(&user)

	authResponse, err := utils.GenerateAuthResponse(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, nil, errors.New("failed to generate auth response")
	}

	// Converter para UserResponse
	userResponse := &models.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Phone:        user.Phone,
		Role:         user.Role,
		IsVerified:   user.IsVerified,
		IsActive:     user.IsActive,
		ProfileImage: user.ProfileImage,
		LastLoginAt:  user.LastLoginAt,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return userResponse, authResponse, nil
}
