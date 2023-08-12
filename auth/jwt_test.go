package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	id := 123
	lifespan := 1
	secretKey := "my-secret-key"

	token, err := GenerateToken(id, lifespan, secretKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Parse the token to verify its claims
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, true, claims["authorized"])
	assert.Equal(t, float64(id), claims["client_id"])
	expClaim, ok := claims["exp"].(float64)
	assert.True(t, ok)
	assert.InDelta(t, float64(time.Now().Unix()+3600), expClaim, 5) // Within 5 seconds tolerance
}
