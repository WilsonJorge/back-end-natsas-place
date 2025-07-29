package auth

import (
	"errors"
	"time"

	"natsas-place/models"
	"natsas-place/repository"
	"natsas-place/utils"
)

// RegisterUser registra um novo usuário
func RegisterUser(input models.UserRegistrationInput) (*models.UserResponse, *utils.AuthResponse, error) {
	// Verificar se o email já está em uso
	existingUser, err := repository.GetUserByEmail(input.Email)
	if err == nil && existingUser.ID > 0 {
		return nil, nil, errors.New("email already in use")
	}

	// Hash da senha
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, nil, errors.New("failed to hash password")
	}

	// Parse da data de nascimento
	dateOfBirth, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		return nil, nil, errors.New("invalid date of birth format")
	}

	// Criar usuário
	user := models.User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    hashedPassword,
		Phone:       input.Phone,
		DateOfBirth: dateOfBirth,
		Gender:      input.Gender,
		Nationality: input.Nationality,
		Language:    input.Language,
		Role:        "USER",
		IsActive:    true,
	}

	// Salvar usuário no banco
	createdUser, err := repository.CreateUser(user)
	if err != nil {
		return nil, nil, errors.New("failed to create user")
	}

	// Gerar resposta de autenticação
	authResponse, err := utils.GenerateAuthResponse(createdUser.ID, createdUser.Email, createdUser.Role)
	if err != nil {
		return nil, nil, errors.New("failed to generate auth response")
	}

	// Converter para UserResponse
	userResponse := &models.UserResponse{
		ID:              createdUser.ID,
		Name:            createdUser.Name,
		Email:           createdUser.Email,
		Phone:           createdUser.Phone,
		Role:            createdUser.Role,
		IsVerified:      createdUser.IsVerified,
		IsActive:        createdUser.IsActive,
		ProfileImage:    createdUser.ProfileImage,
		Gender:          createdUser.Gender,
		Nationality:     createdUser.Nationality,
		Language:        createdUser.Language,
		EmailVerifiedAt: createdUser.EmailVerifiedAt,
		LastLoginAt:     createdUser.LastLoginAt,
		CreatedAt:       createdUser.CreatedAt,
		UpdatedAt:       createdUser.UpdatedAt,
	}

	// Enviar email de verificação (opcional)
	// go func() {
	// 	token, err := utils.GenerateRandomToken()
	// 	if err == nil {
	// 		// Atualizar token no banco
	// 		repository.UpdateUserVerificationToken(createdUser.ID, token)
	// 		// Enviar email (implementar)
	// 		// SendVerificationEmail(createdUser.Email, token)
	// 	}
	// }()

	return userResponse, authResponse, nil
}

// LoginUser autentica um usuário
func LoginUser(input models.UserLoginInput) (*models.UserResponse, *utils.AuthResponse, error) {
	// Buscar usuário pelo email
	user, err := repository.GetUserByEmail(input.Email)
	if err != nil {
		return nil, nil, errors.New("invalid email or password")
	}

	// Verificar se o usuário está ativo
	if !user.IsActive {
		return nil, nil, errors.New("user account is inactive")
	}

	// Verificar senha
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return nil, nil, errors.New("invalid email or password")
	}

	// Atualizar último login
	now := time.Now()
	user.LastLoginAt = &now
	repository.UpdateUser(user)

	// Gerar resposta de autenticação
	authResponse, err := utils.GenerateAuthResponse(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, nil, errors.New("failed to generate auth response")
	}

	// Converter para UserResponse
	userResponse := &models.UserResponse{
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

	return userResponse, authResponse, nil
}
