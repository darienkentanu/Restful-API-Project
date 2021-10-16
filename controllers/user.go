package controllers

import (
	"altastore/lib/database"
	"altastore/middlewares"
	"altastore/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type M map[string]interface{}

func LoginUsersController(c echo.Context) error {
	userLogin := models.UserLogin{}
	c.Bind(&userLogin)

	user, err := database.GetUserByEmail(userLogin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect email")
	}

	check := CheckPasswordHash(userLogin.Password, user.Password)
	if !check {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect password")
	}

	var newToken string
	newToken, err = middlewares.CreateToken(int(user.ID), user.Role)
	if err != nil {
		fmt.Println("gagal bikin token")
		return c.String(http.StatusBadRequest, "Cannot login")
	}

	user.Token = newToken
	user, err = database.UpdateTokenUser(int(user.ID), newToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Cannot add token")
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   user,
	})
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUsersController(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	var err error
	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error in password hash")
	}

	user, err = database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var cart models.Cart
	cart.UserID = user.ID
	err = database.CreateCart(cart)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   user,
	})
}

func GetAllUsersController(c echo.Context) error {
	users, err := database.GetAllUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   users,
	})
}

func UpdateUserController(c echo.Context) error {
	var newUser models.User

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	if id != middlewares.CurrentLoginUser(c) {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unauthorized")
	}

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	user, err := database.UpdateUser(id, newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   user,
	})
}
