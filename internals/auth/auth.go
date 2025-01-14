package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a struct to represent the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Secret key used to sign the JWT tokens
var jwtKey = []byte("your_secret_key")

// GenerateToken generates a JWT token for a given username
func GenerateToken(username string) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the username if valid
func ValidateToken(tokenString string) (string, error) {
	// Parse the token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", errors.New("invalid token signature")
		}
		return "", errors.New("invalid token")
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.Username, nil
}
