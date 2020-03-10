package middleware

import (
	"github.com/gin-gonic/gin"
	"goerhubApi/constraint/e"
	"goerhubApi/helpers"
	"log"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			e.AbortError(c, 401, e.ErrEmptyAuthHeader)
			return
		}
		n := token[7:]
		log.Printf("%s\n", token)
		log.Printf("%s\n", n)
		claims, err := helpers.ParseToken(n)
		if err != nil {
			e.AbortError(c, 401, e.ErrInvalidSigningAlgorithm)
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			e.AbortError(c, 401, e.ErrExpiredToken)
			return
		}
		c.Set("authUserId", claims.UserId)
		c.Next()
	}
}
