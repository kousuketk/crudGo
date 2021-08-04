package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
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

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}

func Authorization(c *gin.Context) int {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user"]
	if str, ok := userId.(string); ok {
		userId, _ = strconv.Atoi(str)
	} else {
		fmt.Println("cast failed.")
	}

	return userId.(int)
}
