package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"natsas-place/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

// HashPassword gera hash da senha usando bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash verifica se a senha corresponde ao hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT gera um token JWT para o usuário
func GenerateJWT(userID uint, email, role string) (string, error) {
	cfg := config.GetConfig()

	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cfg.JWTExpirationHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "natsas-place",
			Subject:   fmt.Sprintf("user:%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

// GenerateRefreshToken gera um refresh token
func GenerateRefreshToken(userID uint) (string, error) {
	cfg := config.GetConfig()

	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)), // 30 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "natsas-place",
			Subject:   fmt.Sprintf("refresh:%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

// ValidateJWT valida um token JWT
func ValidateJWT(tokenString string) (*JWTClaims, error) {
	cfg := config.GetConfig()

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GenerateRandomToken gera um token aleatório para reset de senha
func GenerateRandomToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateAuthResponse gera uma resposta de autenticação completa
func GenerateAuthResponse(userID uint, email, role string) (*AuthResponse, error) {
	cfg := config.GetConfig()

	accessToken, err := GenerateJWT(userID, email, role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(cfg.JWTExpirationHours * 3600), // Convert hours to seconds
	}, nil
}

// ExtractUserFromToken extrai informações do usuário do token
func ExtractUserFromToken(tokenString string) (uint, string, string, error) {
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return 0, "", "", err
	}

	return claims.UserID, claims.Email, claims.Role, nil
}

// IsTokenExpired verifica se o token está expirado
func IsTokenExpired(tokenString string) bool {
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return true
	}

	return claims.ExpiresAt.Time.Before(time.Now())
}

// RefreshAccessToken renova o access token usando refresh token
func RefreshAccessToken(refreshToken string) (*AuthResponse, error) {
	claims, err := ValidateJWT(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Verificar se é um refresh token
	if claims.Subject == "" || claims.Subject[:7] != "refresh" {
		return nil, errors.New("invalid refresh token")
	}

	// Aqui você deveria buscar o usuário do banco de dados para obter email e role atualizados
	// Por simplicidade, vamos usar os dados do token
	return GenerateAuthResponse(claims.UserID, claims.Email, claims.Role)
}

// ValidatePasswordStrength valida a força da senha
func ValidatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' || char == '&' || char == '*':
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character (!@#$%^&*)")
	}

	return nil
}
