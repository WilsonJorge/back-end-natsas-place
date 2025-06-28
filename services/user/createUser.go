package user

import (
	"fmt"
	"natsas-place/models"
	"natsas-place/repository"
)

func CreateUserService(userData models.User) (models.User, error) {

	userExited, err := repository.GetUserSpecification("email", userData.Email)
	if err != nil {
		return models.User{}, fmt.Errorf("THIS USER ALREADY EXISTED")
	}

	if userExited.Email == userData.Email {
		return models.User{}, fmt.Errorf("THIS USER ALREADY EXISTED")
	}

	err = repository.CreateUser(userData)
	if err != nil {
		return models.User{}, fmt.Errorf("USER CREATE FAILED")
	}
	return userData, nil
}
