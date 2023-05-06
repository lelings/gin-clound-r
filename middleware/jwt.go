package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func Auth(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户未登录",
		})
		c.Abort()
		return
	}
	tokenString := authorization[7:]
	claims, err := ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "用户未登录",
		})
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}
