package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/Ethea2/Distributed-Fault-Tolerance/services/common/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Role     string `json:"role,omitempty"`
}

func SignToken(user models.User) (string, error) {
	godotenv.Load()

	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	finalToken, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return finalToken, nil
}

func DecodeToken(tokenstring string) models.User {
	godotenv.Load()
	var custom CustomClaims

	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenstring, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println(err.Error())
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		custom = *claims
	}

	return models.User{
		ID:       custom.ID,
		Username: custom.Username,
		Role:     custom.Role,
	}
}
