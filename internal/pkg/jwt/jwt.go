package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

const key = "r4nd0ms3cr3tk3y"

func GenerateNewTokenString() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	ss, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return ss, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return token, nil
}
