package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Payload interface{}
	jwt.RegisteredClaims
}

type JWTParams struct {
	Payload    interface{}
	SecretKey  string
	ExpireTime int
	issuer     string
}

func GenerateJwtToken(params JWTParams) (string, *time.Time, error) {
	expirationTime := time.Now().Add(time.Duration(params.ExpireTime) * time.Hour)
	claims := &CustomClaims{
		Payload: params.Payload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    params.issuer,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := t.SignedString([]byte(params.SecretKey))
	if err != nil {
		return "", nil, err
	}
	return signed, &expirationTime, nil
}

func VerifiJwtToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token's signing method is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret signing key
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the token and claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
