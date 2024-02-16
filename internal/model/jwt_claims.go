package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"user_id"`
}

func GenerateToken(guid string, ttlToken time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttlToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: guid,
	})

	return token.SignedString([]byte(viper.GetString("signingKey")))
}
