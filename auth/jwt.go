package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateToken generates a JWT token with provided parameters.
func GenerateToken(id, lifespan int, secretKey string) (string, error) {
	// Create a map to hold the JWT claims.
	claims := jwt.MapClaims{}

	// Set the "authorized" claim to true.
	claims["authorized"] = true

	// Set the "client_id" claim to the provided account ID.
	claims["client_id"] = id

	// If a non-zero lifespan is provided, set the "exp" claim to the current time plus the specified duration.
	if lifespan != 0 {
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()
	}

	// Create a new JWT token using the HS256 signing method and the claims map.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the provided secret key and return the signed token string along with any error.
	return token.SignedString([]byte(secretKey))
}
