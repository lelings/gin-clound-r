package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECERT = []byte("rust-go")

type Claims struct {
	Id       int
	Username string
	jwt.StandardClaims
}

func SpawnToken(id int, username string) string {
	claims := Claims{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 14).Unix(),
			Issuer:    "lelings",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECERT)
	if err != nil {
		log.Fatal("[SpawnToken error]:", err)
	}
	return tokenString
}

func ParseToken(tokenString string) (*Claims, error) {
	cliams := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, cliams, func(t *jwt.Token) (interface{}, error) {
		return SECERT, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token 不合法")
	}
	return cliams, nil
}
