package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joshua468/voting-app/internal/config"
	"github.com/joshua468/voting-app/internal/models"
)

func GenerateToken(userID uint, role string) (string, error) {
	claims := &models.JWTClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	secret := config.GetJWTSecret()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString string) (*models.JWTClaims, error) {
	secret := config.GetJWTSecret()

	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
