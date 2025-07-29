package repository

import (
	"natsas-place/database"
	"natsas-place/models"
)

func CreateUser(user models.User) (models.User, error) {
	err := database.Instance.Create(&user).Error
	return user, err
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.Instance.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id any) (models.User, error) {
	var user models.User
	err := database.Instance.First(&user, "id = ?", id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := database.Instance.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(user models.User) (models.User, error) {
	err := database.Instance.Save(&user).Error
	return user, err
}

func DeleteUserById(id string) error {
	return database.Instance.Delete(&models.User{}, "id = ?", id).Error
}
