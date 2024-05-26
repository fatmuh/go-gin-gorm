package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTKey))

	if err != nil {
		return "", fmt.Errorf("Generating JWT Token Failed: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(secretJWTKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Parse JWT Token Failed: %w", err)
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("Invalid JWT Token")
	}
	return claims["sub"], nil
}
