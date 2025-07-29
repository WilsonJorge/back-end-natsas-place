package user

import (
	"fmt"
	"natsas-place/database"
	"natsas-place/models"
)

func CreateUserService(userData models.User) (models.User, error) {
	var userExited models.User
	err := database.Instance.Where("email = ?", userData.Email).First(&userExited).Error
	if err != nil {
		return models.User{}, fmt.Errorf("THIS USER ALREADY EXISTED")
	}

	if userExited.Email == userData.Email {
		return models.User{}, fmt.Errorf("THIS USER ALREADY EXISTED")
	}

	err = database.Instance.Create(&userData).Error
	if err != nil {
		return models.User{}, fmt.Errorf("USER CREATE FAILED")
	}
	return userData, nil
}

func UpdateUser(userData models.User) (models.User, error) {
	var user models.User
	err := database.Instance.Where("id = ?", userData.ID).First(&user).Error
	if err != nil {
		return user, err
	}
	user.Name = userData.Name
	user.Email = userData.Email
	user.Phone = userData.Phone

	if err != nil {
		return user, err
	}

	user.ProfileImage = userData.ProfileImage

	err = database.Instance.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(id string) error {
	err := database.Instance.Delete(&models.User{}, id).Error
	if err != nil {
		return fmt.Errorf("ERROR WHEN DELETE USER %w", err)
	}

	return nil
}

func GetUserById(id string) (models.User, error) {
	var user models.User
	err := database.Instance.Where("id = ?", id).First(&user).Error

	if err != nil {
		return models.User{}, fmt.Errorf("ERROR WHEN GET USER %w", err)
	}

	return user, nil
}

func GetAllUsersService() ([]models.User, error) {
	var users []models.User
	err := database.Instance.Find(&users).Error
	if err != nil {
		return []models.User{}, fmt.Errorf("ERROR WHEN GET USER %w", err)
	}
	return users, nil
}
