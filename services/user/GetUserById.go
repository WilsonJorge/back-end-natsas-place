package user

import (
	"fmt"
	"natsas-place/models"
	"natsas-place/repository"
)

func GetUserById(id string) (models.User, error) {

	user, err := repository.GetUserById(id)

	if err != nil {
		return models.User{}, fmt.Errorf("ERROR WHEN GET USER %w", err)
	}

	return user, nil
}
