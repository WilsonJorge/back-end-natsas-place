package repository

import (
	"natsas-place/database"
	"natsas-place/models"
)

func CreateUser(user models.User) error {
	return database.Instance.Create(user).Error
}

func GetAllUsers() ([]models.User, error) {

	var users []models.User
	err := database.Instance.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id string) (models.User, error) {
	var user models.User
	err := database.Instance.First(&user, "id = ?", id).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateUser(user *models.User) error {
	return database.Instance.Save(user).Error
}

func DeleteUserById(id string) error {
	return database.Instance.Delete(&models.User{}, "id = ?", id).Error
}

//specifications

func GetUserSpecification(where, what string) (models.User, error) {
	var user models.User
	err := database.Instance.Model(&user).Where(where, what).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
