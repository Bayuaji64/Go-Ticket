package helper

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userId int64) (string, error) {

	jwtKey := "rahasia"

	if jwtKey == "" {
		return "", errors.New("JWT_KEY is not set in the environment variables")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(jwtKey))
}

func VerifyToken(token string) (int64, error) {

	jwtKey := "rahasia"
	if jwtKey == "" {
		return 0, errors.New("JWT_KEY is not set in the environment variables")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtKey), nil
	})

	if err != nil {

		log.Printf("Error parsing token: %v", err)
		return 0, fmt.Errorf("could not parse token: %v", err)
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)

	userId := int64(claims["userId"].(float64))
	return userId, err
}
