package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateTokenForUser(email string, uuid uuid.UUID) ([]byte, error) {
	rsaKey, err := GetRsaPrivateKey()
	if err != nil {
		return nil, err
	}

	jwkKey, err := GetJwk()
	if err != nil {
		return nil, err
	}

	// 7 days
	expiry := time.Now().Add(time.Hour * 24 * 7).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":   uuid.String(),
		"email": email,
		"iss":   "https://vaults.dev",
		"exp":   expiry,
		"iat":   time.Now().Unix(),
	})
	jwtToken.Header["kid"] = jwkKey.Kid
	jwtToken.Header["alg"] = jwt.SigningMethodRS256.Name

	token, err := jwtToken.SignedString(rsaKey)
	if err != nil {
		return nil, err
	}

	return []byte(token), nil
}
