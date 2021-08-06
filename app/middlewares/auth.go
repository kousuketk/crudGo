package middlewares

import (
	"net/http"
	"os"
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
	secretKey := os.Getenv("SECRETKEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	secretKey := os.Getenv("SECRETKEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}

func Authorization(c *gin.Context) (int, bool) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return -1, false
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user"]
	if str, ok := userId.(string); ok {
		userId, _ = strconv.Atoi(str)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cast failed."})
		return -1, false
	}

	return userId.(int), true
}
