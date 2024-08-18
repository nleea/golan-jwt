package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return "", nil
		})

		if err != nil {
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			fmt.Println(err)
		}

		ctx.Next()
	}
}
