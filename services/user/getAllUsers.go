package user

import (
	"fmt"
	"natsas-place/models"
	"natsas-place/repository"
)

func GetAllUsersService() ([]models.User, error) {
	userData, err := repository.GetAllUsers()
	if err != nil {
		return []models.User{}, fmt.Errorf("ERROR WHEN GET USER %w", err)
	}
	return userData, nil
}
