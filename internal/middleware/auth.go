package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	core_errors "github.com/wavw1/music-tracker/internal/errors"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": core_errors.ErrEmptySecret.Error()})
			c.Abort()
			return
		}

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": core_errors.ErrTokenRequired.Error()})
			c.Abort()
			return
		}

		parts := strings.Fields(tokenString)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": core_errors.ErrTokenFormat.Error()})
			c.Abort()
			return
		}

		tokenString = parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": core_errors.ErrInvalidToken})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		subFloat, ok := claims["sub"].(float64)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID := int64(subFloat)
		c.Set("user_id", userID)

		c.Next()
	}
}
