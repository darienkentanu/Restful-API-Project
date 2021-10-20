package database

import (
	"altastore/config"
	"altastore/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserByEmail(userLogin models.UserLogin) (models.User, error) {
	user := models.User{}
	err := config.InitDB().Where("email = ?", userLogin.Email).First(&user).Error
	if err != nil {
		return models.User{}, echo.NewHTTPError(http.StatusInternalServerError)
	}
	return user, nil
}

func GetUserByID(userID int) (models.User, error) {
	user := models.User{}
	if err := config.InitDB().Where("id = ?", userID).First(&user).Error; err != nil {
		return models.User{}, echo.NewHTTPError(http.StatusInternalServerError)
	}
	return user, nil
}

func UpdateUser(id int, newUser models.User) (models.User, error) {
	var user models.User

	if err := config.InitDB().First(&user, id).Error; err != nil {
		return user, err
	}

	user.Fullname = newUser.Fullname
	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Password = newUser.Password
	user.PhoneNumber = newUser.PhoneNumber
	user.Gender = newUser.Gender
	user.Address = newUser.Address

	if err := config.InitDB().Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateTokenUser(id int, newToken string) (models.User, error) {
	var user models.User
	if err := config.InitDB().First(&user, id).Error; err != nil {
		return user, err
	}

	user.Token = newToken

	if err := config.InitDB().Model(&user).Update("token", newToken).Error; err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	user.Role = "customer"

	if err := config.InitDB().Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := config.InitDB().Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
