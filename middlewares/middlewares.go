package middlewares

import (
	"altastore/constants"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
	SigningKey: []byte(constants.JWT_SECRET),
})

func CreateToken(userId int, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires after 1 hour

	tokenString, err := token.SignedString([]byte(constants.JWT_SECRET))
	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		token := e.Get("user").(*jwt.Token)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			role := claims["role"].(string)
			if role != "admin" {
				return echo.ErrUnauthorized
			}
		}
		return next(e)
	}
}

func CurrentLoginUser(e echo.Context) int {
	token := e.Get("user").(*jwt.Token)
	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"]
		switch userId.(type) {
		case float64:
			return int(userId.(float64))
		default:
			return userId.(int)
		}
	}
	return -1 // invalid user
}
