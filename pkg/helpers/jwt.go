package helpers

import (
	"errors"
	"final_project_3/infrastructure/config"
	"strings"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id int, email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"role": role,
	})
	tokenString, err := token.SignedString([]byte(config.GetAppConfig().JWTSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResp := errors.New("invalid authorization token")
	getToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(getToken, "Bearer")

	if !bearer{
		return nil, errResp
	}

	tokenString := strings.Split(getToken, " ")[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.GetAppConfig().JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}