package middlewares

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(personId int, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["personId"] = personId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func CreateTokenBook(personId int, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["personId"] = personId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// func ExtractTokenUserId(e echo.Context) int {
// 	token := e.Get("user").(*jwt.Token)
// 	if token.Valid {
// 		claims := token.Claims.(jwt.MapClaims)
// 		personId := claims["personId"].(int)
// 		return personId
// 	}
// 	return -1 // invalid user
// }
