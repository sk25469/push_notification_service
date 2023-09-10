package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sk25469/push_noti_service/pkg/utils"
)

func CreateJWTToken(username string) (string, error) {
	// Create a new token with claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(utils.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
