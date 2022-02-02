package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/DaniaRepublic/commonSpaceGo/classes"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type authClaims struct {
	jwt.StandardClaims
	PhoneNum   string
	CreateTime string
}

func GenerateToken(u classes.User, ttl int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   u.Username,
			ExpiresAt: time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
		},
		PhoneNum:   u.PhoneNum,
		CreateTime: u.CreateTime,
	})

	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func validateToken(tokenStr string) (string, string, error) {
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", errors.New("invalid token")
	}

	return claims.Subject, claims.PhoneNum, nil
}
