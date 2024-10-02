package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret")

func CreateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["sub"].(string)
		return userID, nil
	}

	return "", fmt.Errorf("invalid token")
}

func main() {
	// Пример использования
	token, err := CreateToken("user123")
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}

	fmt.Println("Generated Token:", token)

	userID, err := ValidateToken(token)
	if err != nil {
		fmt.Println("Error validating token:", err)
		return
	}

	fmt.Println("Valid Token! User ID:", userID)
}
