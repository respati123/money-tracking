package util

import (
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

func VerifiJwtToken(tokenString, secretKey string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.Payload, nil
	}
	return nil, err
}
