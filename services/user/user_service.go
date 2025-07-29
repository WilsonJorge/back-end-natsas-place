package user

import (
	"fmt"
	"natsas-place/models"
	"natsas-place/repository"
)

func CreateUserService(userData models.User) (models.User, error) {
	userExited, err := repository.GetUserByEmail(userData.Email)
	if err != nil {
		return models.User{}, fmt.Errorf("THIS USER ALREADY EXISTED")
	}

	if userExited.Email == userData.Email {
		return models.User{}, fmt.Errorf("THIS USER ALREADY EXISTED")
	}

	_, err = repository.CreateUser(userData)
	if err != nil {
		return models.User{}, fmt.Errorf("USER CREATE FAILED")
	}
	return userData, nil
}

func UpdateUser(userData models.User) (models.User, error) {
	user, err := repository.GetUserByID(userData.ID)
	if err != nil {
		return user, err
	}
	user.Name = userData.Name
	user.Email = userData.Email
	user.Phone = userData.Phone
	user.DateOfBirth = userData.DateOfBirth
	if err != nil {
		return user, err
	}

	user.ProfileImage = userData.ProfileImage

	_, err = repository.UpdateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(id string) error {
	err := repository.DeleteUserById(id)
	if err != nil {
		return fmt.Errorf("ERROR WHEN DELETE USER %w", err)
	}

	return nil
}

func GetUserById(id string) (models.User, error) {
	user, err := repository.GetUserByID(id)

	if err != nil {
		return models.User{}, fmt.Errorf("ERROR WHEN GET USER %w", err)
	}

	return user, nil
}

func GetAllUsersService() ([]models.User, error) {
	userData, err := repository.GetAllUsers()
	if err != nil {
		return []models.User{}, fmt.Errorf("ERROR WHEN GET USER %w", err)
	}
	return userData, nil
}
