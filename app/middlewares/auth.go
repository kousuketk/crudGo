package middlewares

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

func CreateToken(userId string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"user": userId,
		"exp":  time.Now().Add(time.Hour * 1).Unix(), // 有効期限を指定
	}
	var secretKey = "secret" // 任意の文字列4
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
